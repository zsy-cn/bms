package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	stdlog "log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/go-kit/kit/endpoint"
	gkLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"

	gokitConsul "github.com/go-kit/kit/sd/consul"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
)

// ParserHub 增删服务
type ParserHub struct {
	logger    *log.Logger
	db        *gorm.DB
	consulCli gokitConsul.Client
}

// New ...
func New(logger *log.Logger) (parserhub *ParserHub, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}
	consulCli, err := conf.ConnectConsul()
	if err != nil {
		logger.Errorf("connect consul failed: %s", err.Error())
		return
	}

	parserhub = &ParserHub{
		logger:    logger,
		db:        db,
		consulCli: consulCli,
	}
	return
}

// ParseAndSave ...
// 首先将payload进行base64解码为[]byte类型数据,
// 上行原始数据中带有eui信息, 根据devEUI查找该设备的型号, 再根据型号去解析payload.
// 拿到返回值后入库, 原始[]byte值和解码后的字符串一起
func (ph *ParserHub) ParseAndSave(uplinkMsg *protos.ParserHubUplinkMsg) (err error) {
	ph.logger.Debug("parse uplink message in ParseAndSave()")
	// 对上行数据进行base64解码
	data, err := base64.StdEncoding.DecodeString(uplinkMsg.Data)
	if err != nil {
		ph.logger.Errorf("base64 decode uplink message payload failed: %s", err.Error())
		return
	}
	uplinkMsg.FinalData = data
	// 获取可用于此条上行信息的endpoint接口
	reqEndpoint, err := ph.getParserEndpoint(uplinkMsg.DevEUI)
	if err != nil {
		// 错误日志在getParserEndpoint函数体内部打印了
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// 手动调用endpoint, 向parser子服务发起请求
	_, err = reqEndpoint(ctx, uplinkMsg)
	return
}

// getParserEndpoint 获取可用于目标上行信息的endpoint接口, 需要通过consul中的服务名称查询
func (ph *ParserHub) getParserEndpoint(devEUI string) (endpoint endpoint.Endpoint, err error) {
	gkLogger := gkLog.NewNopLogger()
	serviceName, err := ph.getServiceNameForSensor(devEUI)
	// serviceName := "trashcan_lierda_01"
	tags := []string{serviceName}
	passingOnly := true
	//创建实例管理器
	instancer := gokitConsul.NewInstancer(ph.consulCli, gkLogger, serviceName, tags, passingOnly)
	//创建端点管理器， 此管理器根据Factory和监听到的实例创建endPoint并订阅instancer的变化动态更新Factory创建的endPoint
	endpointer := sd.NewEndpointer(instancer, dispatchFactory, gkLogger)
	endpointList, err := endpointer.Endpoints()
	// 如果没有在consul中发现目标服务, 则丢弃该条消息
	if len(endpointList) == 0 {
		ph.logger.Warnf("can not found endpoints for service: %s, uplink message dropped", serviceName)
		err = errors.New("endpoints not found")
		return
	}
	//创建负载均衡器
	balancer := lb.NewRoundRobin(endpointer)
	endpoint, err = balancer.Endpoint()
	return
}

// 根据传感器的devEUI查询其设备类型及型号, 组成服务名称
func (ph *ParserHub) getServiceNameForSensor(devEUI string) (serviceName string, err error) {
	sensorModel := &model.Sensor{}
	whereArgs := map[string]interface{}{
		"dev_eui": devEUI,
	}
	err = ph.db.Where(whereArgs).First(sensorModel).Error
	if err != nil {
		if err.Error() != "record not found" {
			ph.logger.Errorf("query sensor with dev_eui %s failed: %s", devEUI, err.Error())
			return
		}
		return
	}
	deviceID := sensorModel.DeviceSN
	deviceRecord := &model.Device{}
	err = ph.db.First(deviceRecord, deviceID).Error
	if err != nil {
		if err.Error() == "record not found" {
			ph.logger.Warnf("can not find device for sensor dev_eui %s with device_id %d, you should repair it", devEUI, deviceID)
		} else {
			ph.logger.Errorf("query device for sensor dev_eui %s with device_id %d failed: %s", devEUI, deviceID, err.Error())
		}
		return
	}
	deviceTypeRecord := &model.DeviceType{}
	deviceModelRecord := &model.DeviceModel{}
	err = ph.db.Model(deviceRecord).Related(deviceTypeRecord).Related(deviceModelRecord).Error
	if err != nil {
		if err.Error() == "record not found" {
			ph.logger.Warnf("can not find device type or model for device %d, you should repair it", deviceID)
		} else {
			ph.logger.Errorf("query device type and model for device: %d failed: %s", deviceID, err.Error())
		}
		return
	}
	serviceName = deviceTypeRecord.Key + "_" + deviceModelRecord.Name
	return
}

// 通过传入的子服务实例地址instance, 创建对应的请求endPoint
func dispatchFactory(instance string) (endpoint.Endpoint, io.Closer, error) {
	stdlog.Println("making endpoint to parser sub service in factory")
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}

	serviceURL, err := url.Parse(instance)
	if err != nil {
		return nil, nil, err
	}
	serviceURL.Path = "/parse"

	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		serviceURLStr := serviceURL.String()
		stdlog.Printf("request sub service at %s\n", serviceURLStr)

		byteData, err := json.Marshal(request)
		if err != nil {
			panic(err)
		}
		jsonData := bytes.NewBuffer(byteData)

		_, err = http.Post(serviceURLStr, "application/json", jsonData)
		if err != nil {
			stdlog.Printf("request sub service at %s failed: %s\n", serviceURLStr, err.Error())
			return
		}
		return &protos.Empty{}, nil
	}, nil, nil
}
