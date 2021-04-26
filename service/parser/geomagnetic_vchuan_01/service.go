package geomagnetic_vchuan_01

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/geomagnetic"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/parser"
	"github.com/zsy-cn/bms/util/log"
)

// GeomagneticVchuan01 ...
type GeomagneticVchuan01 struct {
	parser.ConsulParserService
	logger *log.Logger
	db     *gorm.DB
}

// New ...
func New(logger *log.Logger) (device *GeomagneticVchuan01, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}
	device = &GeomagneticVchuan01{
		logger: logger,
		db:     db,
	}
	return
}

// Decode ...
// 要实现一个interface, receiver不能是指针???(GeomagneticVchuan01实现api.go->Parser接口)
// uplinkMsg中FinalData已经经过base64解码, 无需重复操作, 直接解析即可
func (device GeomagneticVchuan01) Decode(uplinkMsg *protos.ParserHubUplinkMsg) (err error) {
	device.logger.Debugf("receive uplink mesage in Decode(): %+v", uplinkMsg)
	device.logger.Debug("uplink mesage payload bytes in Decode(): ", uplinkMsg.FinalData)

	uplinkMsgPayload := &model.UplinkMsgPayload{
		Attributes: model.UplinkMsgAttribute{},
	}
	byteData := uplinkMsg.FinalData

	header := byteData[0]
	if header != 0xAA {
		device.logger.Error("unknown package, drop it")
		return
	}
	pktType := byteData[1]

	msgType := geomagnetic.GeomagneticMsgType_Triggered
	if pktType == 0x31 {
		msgType = geomagnetic.GeomagneticMsgType_Triggered
		parser00(byteData, uplinkMsgPayload)
	} else if pktType == 0x32 {
		msgType = geomagnetic.GeomagneticMsgType_Confirm
		parser00(byteData, uplinkMsgPayload)
	} else if pktType == 0x33 {
		msgType = geomagnetic.GeomagneticMsgType_Heartbeat
		parser00(byteData, uplinkMsgPayload)
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

	err = fillModel(geomagneticModel, uplinkMsgPayload)
	if err != nil {
		device.logger.Errorf("fill geomagnetic model failed in Decode(): %s", err.Error())
		return
	}

	err = device.db.Create(geomagneticModel).Error
	if err != nil {
		device.logger.Errorf("create geomagnetic app message record failed in Decode(): %s", err.Error())
		return
	}
	return
}

func fillModel(geomagneticModel *geomagnetic.Geomagnetic, uplinkMsgPayload *model.UplinkMsgPayload) (err error) {
	used, exist := uplinkMsgPayload.Attributes["Used"]
	if exist && used != nil {
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

	return
}
