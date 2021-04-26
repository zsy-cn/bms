package temperature_weichuan_01

import (
	"bytes"
	"encoding/binary"

	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/parser"
	"github.com/zsy-cn/bms/util/log"
)

// TemperatureWeichuan01 ...
type TemperatureWeichuan01 struct {
	parser.ConsulParserService
	logger *log.Logger
	db     *gorm.DB
}

// New ...
func New(logger *log.Logger) (device *TemperatureWeichuan01, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}
	device = &TemperatureWeichuan01{
		logger: logger,
		db:     db,
	}
	return
}

// Decode ...
// 要实现一个interface, receiver不能是指针???(TemperatureWeichuan01实现api.go->Parser接口)
// uplinkMsg中FinalData已经经过base64解码, 无需重复操作, 直接解析即可
func (device TemperatureWeichuan01) Decode(uplinkMsg *protos.ParserHubUplinkMsg) (err error) {
	device.logger.Debugf("receive uplink mesage in Decode(): %+v", uplinkMsg)
	device.logger.Debug("uplink mesage payload bytes in Decode(): ", uplinkMsg.FinalData)

	uplinkMsgPayload := &model.UplinkMsgPayload{
		Attributes: model.UplinkMsgAttribute{},
	}
	byteData := uplinkMsg.FinalData

	////////////////////////////////////////////////////////////////////////////
	// 解析过程
	// 注意: 高字节在前
	rawTemperature := [2]byte{byteData[0], byteData[1]}
	parser.TemperatureDecoder(uplinkMsgPayload.Attributes, rawTemperature, "温度")

	buf := bytes.NewBuffer(byteData[2:4])
	var intHumidity int16
	binary.Read(buf, binary.BigEndian, &intHumidity)

	humidityAttr := &model.UplinkMsgPayloadField{
		DisplayName: "湿度",
		Unit:        "%",
		Value:       float64(intHumidity) / 100,
	}
	uplinkMsgPayload.Attributes["Humidity"] = humidityAttr
	rawVoltage := [1]byte{byteData[4]}
	parser.VoltageDecoder(uplinkMsgPayload.Attributes, rawVoltage, "电压")

	////////////////////////////////////////////////////////////////////
	// 解析完成后入库
	baseAppModel := &model.App{}
	err = parser.FillBaseAppModel(device.db, uplinkMsg, uplinkMsgPayload, baseAppModel)
	if err != nil {
		// 错误日志在FillBaseAppModel函数内部打印过了
		return
	}
	Model := &model.Temperature{
		App:         *baseAppModel, // 引用类型
		Temperature: uplinkMsgPayload.Attributes["Temperature"].Value,
		Humidity:    uplinkMsgPayload.Attributes["Humidity"].Value,
		Voltage:     uplinkMsgPayload.Attributes["Voltage"].Value,
	}

	err = device.db.Create(Model).Error
	if err != nil {
		device.logger.Errorf("create temperature app message record failed in Decode(): %s", err.Error())
		return
	}
	return
}
