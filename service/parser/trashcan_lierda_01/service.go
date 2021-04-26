package trashcan_lierda_01

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/parser"
	"github.com/zsy-cn/bms/trashcan"
	"github.com/zsy-cn/bms/util/log"
)

// TrashcanLierda01 ...
type TrashcanLierda01 struct {
	parser.ConsulParserService
	logger *log.Logger
	db     *gorm.DB
}

// New ...
func New(logger *log.Logger) (device *TrashcanLierda01, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}
	device = &TrashcanLierda01{
		logger: logger,
		db:     db,
	}
	return
}

// Decode ...
// 要实现一个interface, receiver不能是指针???(在这里是TrashcanLierda01实现api.go->Parser接口)
// uplinkMsg中FinalData已经经过base64解码, 无需重复操作, 直接解析即可
func (device TrashcanLierda01) Decode(uplinkMsg *protos.ParserHubUplinkMsg) (err error) {
	device.logger.Debugf("receive uplink message in Decode(): %+v", uplinkMsg)

	uplinkMsgPayload := &model.UplinkMsgPayload{
		Attributes: model.UplinkMsgAttribute{},
	}
	byteData := uplinkMsg.FinalData
	if byteData[0] != 0x80 && byteData[len(byteData)-1] != 0x81 {
		device.logger.Warn("unknown package type, drop it")
		return
	}
	pktType := byteData[1]

	var msgType trashcan.TrashcanMsgType
	if pktType == 0x03 {
		msgType = trashcan.TrashcanMsgType_Register
		parser03(byteData, uplinkMsgPayload)
	} else if pktType == 0x31 {
		msgType = trashcan.TrashcanMsgType_Status
		parser31(byteData, uplinkMsgPayload)
	} else if pktType == 0x32 {
		msgType = trashcan.TrashcanMsgType_Heartbeat
		parser32(byteData, uplinkMsgPayload)
	} else if pktType == 0x33 {
		msgType = trashcan.TrashcanMsgType_Position
		parser33(byteData, uplinkMsgPayload)
	} else if pktType == 0x34 {
		msgType = trashcan.TrashcanMsgType_Status2
		parser34(byteData, uplinkMsgPayload)
	} else if pktType == 0x17 {
		msgType = trashcan.TrashcanMsgType_Reply
		parser17(byteData, uplinkMsgPayload)
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
	trashcanModel := &trashcan.Trashcan{
		App:     *baseAppModel, // 引用类型
		MsgType: msgType,
	}

	err = fillModel(trashcanModel, uplinkMsgPayload)
	if err != nil {
		device.logger.Errorf("fill trashcan model failed in Decode(): %s", err.Error())
		return
	}

	err = device.db.Create(trashcanModel).Error
	if err != nil {
		device.logger.Errorf("create trashcan app message record failed in Decode(): %s", err.Error())
		return
	}

	// 创建异常通知
	err = parser.ReportTrashcanNotification(device.db, trashcanModel)
	if err != nil {
		device.logger.Errorf("try to report trashcan notification failed in Decode(): %s", err.Error())
		return
	}
	return
}

func fillModel(trashcanModel *trashcan.Trashcan, uplinkMsgPayload *model.UplinkMsgPayload) (err error) {
	distance, exist := uplinkMsgPayload.Attributes["Distance"]
	if exist && distance != nil {
		trashcanModel.Percent = distance.Value
	}

	return
}
