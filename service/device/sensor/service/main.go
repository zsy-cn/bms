package service

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/util/log"
)

// DeviceSensor 增删服务
type DeviceSensor struct {
	logger        *log.Logger
	db            *gorm.DB
	loraclientCli protos.LoraclientServiceClient
}

// New ...
func New(logger *log.Logger) (device *DeviceSensor, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}
	address := viper.GetString("loraclient-addr")
	loraclientConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("connect loraclient failed: %s", err.Error())
		return
	}
	loraclientCli := protos.NewLoraclientServiceClient(loraclientConn)

	device = &DeviceSensor{
		logger:        logger,
		db:            db,
		loraclientCli: loraclientCli,
	}
	return
}

// Add ...
func (ds *DeviceSensor) Add(sensorPb *protos.DeviceSensor) (err error) {
	ds.logger.Debugf("add sensor device: %s in Add()", sensorPb.Name)
	sensorModel := &model.Sensor{}
	err = pb2Model(sensorPb, sensorModel)
	if err != nil {
		ds.logger.Errorf("transform sensor: %s protobuf object to model failed in add: %s", sensorPb.Name, err.Error())
		return
	}

	loraclientSensorPb, err := ds.makeLoraclientSensor(sensorPb)
	if err != nil {
		// 错误日志在makeLoraclientSensor打印过了
		return
	}

	tx := ds.db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	err = tx.Create(sensorModel).Error
	if err != nil {
		ds.logger.Errorf("insert sensor: %s into database failed: %s", sensorPb.Name, err.Error())
		return
	}

	// 通过loraclient创建传感器对象
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err = ds.loraclientCli.AddSensor(ctx, loraclientSensorPb)
	if err != nil {
		ds.logger.Errorf("create sensor: %s by loraclient failed: %s", sensorPb.Name, err.Error())
		return
	}
	return
}

// Update ...
func (ds *DeviceSensor) Update(sensorPb *protos.DeviceSensor) (err error) {
	record := &model.Sensor{}
	err = ds.db.Where(&model.Sensor{DeviceSN: sensorPb.DeviceSN}).First(record).Error
	if err != nil {
		ds.logger.Errorf("find sensor: %s failed in update: %s", sensorPb.Name, err.Error())
		return
	}

	sensorMap := map[string]interface{}{}
	err = pb2Map(sensorPb, sensorMap)
	if err != nil {
		ds.logger.Errorf("transform sensor: %s protobuf object to model failed in add: %s", sensorPb.Name, err.Error())
		return
	}

	loraclientSensorPb, err := ds.makeLoraclientSensor(sensorPb)
	if err != nil {
		// 错误日志在makeLoraclientSensor打印过了
		return
	}

	tx := ds.db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	err = tx.Model(record).Update(sensorMap).Error
	if err != nil {
		ds.logger.Errorf("update sensor: %s failed in update: %s", sensorPb.Name, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err = ds.loraclientCli.UpdateSensor(ctx, loraclientSensorPb)
	if err != nil {
		ds.logger.Errorf("update sensor: %s failed from loraclient in update: %s", sensorPb.Name, err.Error())
		return
	}
	return
}

// Delete ...
func (ds *DeviceSensor) Delete(sensorPb *protos.DeviceSensor) (err error) {
	ds.logger.Debugf("delete sensor: %s(device_id) in Delete()", sensorPb.DeviceSN)
	record := &model.Sensor{}
	err = ds.db.Where(&model.Sensor{DeviceSN: sensorPb.DeviceSN}).First(record).Error
	if err != nil {
		ds.logger.Errorf("find sensor: %s failed in Delete(): %s", sensorPb.Name, err.Error())
		return
	}

	tx := ds.db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	err = tx.Delete(record).Error
	if err != nil {
		ds.logger.Errorf("delete sensor: %s failed in Delete(): %s", sensorPb.Name, err.Error())
		return
	}

	loraclientSensorPb := &protos.LoraclientSensor{
		DevEUI: record.DevEUI,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err = ds.loraclientCli.DeleteSensor(ctx, loraclientSensorPb)
	if err != nil {
		ds.logger.Errorf("delete sensor by loraclient: %s failed in Delete(): %s", sensorPb.Name, err.Error())
		return
	}
	return
}

func (ds *DeviceSensor) makeLoraclientSensor(sensorPb *protos.DeviceSensor) (loraclientSensorPb *protos.LoraclientSensor, err error) {
	// 构建LoraclientSensor对象
	customerModel := &model.Customer{}
	err = ds.db.First(customerModel, sensorPb.CustomerID).Error
	if err != nil {
		ds.logger.Errorf("get the customer: %d failed: %s", sensorPb.CustomerID, err.Error())
		return
	}
	sensorTypeModel := &model.DeviceType{}
	sensorTypeWhereArgs := map[string]interface{}{
		"id":        sensorPb.DeviceTypeID,
		"is_sensor": "true",
	}
	err = ds.db.Where(sensorTypeWhereArgs).First(sensorTypeModel).Error
	if err != nil {
		ds.logger.Errorf("get the device type: %d failed: %s", sensorPb.DeviceTypeID, err.Error())
		return
	}
	loraclientSensorPb = &protos.LoraclientSensor{
		OrgDisplayName: customerModel.Title, // 搜索所属org时, 所需的字段是DisplayName
		Type:           sensorTypeModel.Key,
	}
	err = pb2LoraSensorPb(sensorPb, loraclientSensorPb)
	if err != nil {
		ds.logger.Errorf("transform sensor: %s protobuf object to loraclient sensor object failed: %s", sensorPb.Name, err.Error())
		return
	}
	return
}

func pb2Model(pb *protos.DeviceSensor, model *model.Sensor) (err error) {
	model.DeviceSN = pb.DeviceSN
	model.DevEUI = pb.DevEUI
	model.AppEUI = pb.AppEUI
	model.AppKey = pb.AppKey
	model.Freq = pb.Freq
	return
}

func pb2Map(pb *protos.DeviceSensor, theMap map[string]interface{}) (err error) {
	theMap["name"] = pb.Name
	theMap["device_sn"] = pb.DeviceSN

	theMap["dev_eui"] = pb.DevEUI
	theMap["app_eui"] = pb.AppEUI
	theMap["app_key"] = pb.AppKey
	theMap["freq"] = pb.Freq
	return
}

func pb2LoraSensorPb(pb *protos.DeviceSensor, loraclientSensor *protos.LoraclientSensor) (err error) {
	loraclientSensor.Name = pb.Name
	loraclientSensor.DevEUI = pb.DevEUI
	loraclientSensor.AppKey = pb.AppKey
	return
}
