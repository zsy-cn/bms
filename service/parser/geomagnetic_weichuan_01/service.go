package geomagnetic_weichuan_01

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/geomagnetic"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/parser"
	"github.com/zsy-cn/bms/util/log"
)

// GeomagneticWeichuan01 ...
type GeomagneticWeichuan01 struct {
	parser.ConsulParserService
	logger *log.Logger
	db     *gorm.DB
}

// New ...
func New(logger *log.Logger) (device *GeomagneticWeichuan01, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}
	device = &GeomagneticWeichuan01{
		logger: logger,
		db:     db,
	}
	return
}

// Decode ...
// 要实现一个interface, receiver不能是指针???(GeomagneticWeichuan01实现api.go->Parser接口)
// uplinkMsg中FinalData已经经过base64解码, 无需重复操作, 直接解析即可
func (device GeomagneticWeichuan01) Decode(uplinkMsg *protos.ParserHubUplinkMsg) (err error) {
	device.logger.Debugf("receive uplink mesage in Decode(): %+v", uplinkMsg)
	device.logger.Debug("uplink mesage payload bytes in Decode(): ", uplinkMsg.FinalData)

	uplinkMsgPayload := &model.UplinkMsgPayload{
		Attributes: model.UplinkMsgAttribute{},
	}
	byteData := uplinkMsg.FinalData
	pktType := byteData[0] >> 4
	direction := byteData[0] & 0x01

	// // 0为上行, 1为下行
	if direction != 0 {
		device.logger.Info("current message's direction is 1, for downlink, drop it")
		return
	}
	msgType := geomagnetic.GeomagneticMsgType_Heartbeat
	if pktType == 0x01 {
		msgType = geomagnetic.GeomagneticMsgType_Config
		parser01(byteData, uplinkMsgPayload)
	} else if pktType == 0x02 {
		msgType = geomagnetic.GeomagneticMsgType_Status
		parser02(byteData, uplinkMsgPayload)
	} else if pktType == 0x03 {
		msgType = geomagnetic.GeomagneticMsgType_Heartbeat
		parser03(byteData, uplinkMsgPayload)
	} else if pktType == 0x04 {
		msgType = geomagnetic.GeomagneticMsgType_Triggered
		parser04(byteData, uplinkMsgPayload)
	} else if pktType == 0x06 {
		// 该包是confirm 包, 地磁收到该包后, 需要先恢复 ACK, 然后再重启.
		// 此类型的包不携带任何有效信息
		msgType = geomagnetic.GeomagneticMsgType_Confirm
	} else if pktType == 0x07 {
		// 该包是confirm 包, 地磁收到该包后, 需要先恢复 ACK, 然后再重启.
		// 此类型的包不携带任何有效信息
		msgType = geomagnetic.GeomagneticMsgType_XYZ
		parser07(byteData, uplinkMsgPayload)
	} else if pktType == 0x08 {
		msgType = geomagnetic.GeomagneticMsgType_Heartbeat_New
		parser08(byteData, uplinkMsgPayload)
	} else if pktType == 0x0a {
		msgType = geomagnetic.GeomagneticMsgType_Triggered_New
		parser0a(byteData, uplinkMsgPayload)
	} else if pktType == 0x09 {
		msgType = geomagnetic.GeomagneticMsgType_Error
		parser09(byteData, uplinkMsgPayload)
	} else {
		device.logger.Warn("unknown package type, drop it")
		return
	}
	// 解析完成后入库
	baseAppModel := &model.App{}
	err = parser.FillBaseAppModel(device.db, uplinkMsg, uplinkMsgPayload, baseAppModel)
	if err != nil {
		// 错误日志在FillBaseAppModel函数内部打印过了
		return
	}
	geomagneticModel := &geomagnetic.Geomagnetic{
		App:     *baseAppModel, // 引用类型
		MsgType: msgType,
	}

	// 如下消息体类型才会包含停车信息
	if msgType == geomagnetic.GeomagneticMsgType_Status ||
		msgType == geomagnetic.GeomagneticMsgType_Triggered ||
		msgType == geomagnetic.GeomagneticMsgType_Triggered_New ||
		msgType == geomagnetic.GeomagneticMsgType_Heartbeat ||
		msgType == geomagnetic.GeomagneticMsgType_Heartbeat_New {
		used := uplinkMsgPayload.Attributes["Used"]
		if used == nil {
			device.logger.Warn("uplink message payload doesn't have 'Used' field, do not save it")
		}
		if uplinkMsgPayload.Attributes["Used"].Value == 1 {
			geomagneticModel.Value = "1"
		} else if uplinkMsgPayload.Attributes["Used"].Value == 0 {
			geomagneticModel.Value = "0"
		}
	}

	voltage, exist := uplinkMsgPayload.Attributes["Voltage"]
	if exist && voltage != nil {
		geomagneticModel.Voltage = voltage.Value
	}

	xAxis, exist := uplinkMsgPayload.Attributes["XAxis"]
	if exist && xAxis != nil {
		geomagneticModel.XAxis = xAxis.Value
	}
	yAxis, exist := uplinkMsgPayload.Attributes["YAxis"]
	if exist && yAxis != nil {
		geomagneticModel.YAxis = yAxis.Value
	}
	zAxis, exist := uplinkMsgPayload.Attributes["ZAxis"]
	if exist && zAxis != nil {
		geomagneticModel.ZAxis = zAxis.Value
	}

	err = device.db.Create(geomagneticModel).Error
	if err != nil {
		device.logger.Errorf("create geomagnetic app message record failed in Decode(): %s", err.Error())
		return
	}

	device.logger.Debug("had save to database")

	return
}
