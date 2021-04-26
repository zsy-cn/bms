package manhole_cover_gti_01

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/manhole_cover"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/parser"
	"github.com/zsy-cn/bms/util/log"
)

// ManholeCoverGti01 ...
type ManholeCoverGti01 struct {
	parser.ConsulParserService
	logger *log.Logger
	db     *gorm.DB
}

// New ...
func New(logger *log.Logger) (device *ManholeCoverGti01, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}
	device = &ManholeCoverGti01{
		logger: logger,
		db:     db,
	}
	return
}

// Decode ...
// 要实现一个interface, receiver不能是指针???(在这里是ManholeCoverGti01实现api.go->Parser接口)
// uplinkMsg中FinalData已经经过base64解码, 无需重复操作, 直接解析即可
func (device ManholeCoverGti01) Decode(uplinkMsg *protos.ParserHubUplinkMsg) (err error) {
	device.logger.Debugf("receive uplink message in Decode(): %+v", uplinkMsg)
	uplinkMsgPayload := &model.UplinkMsgPayload{
		Attributes: model.UplinkMsgAttribute{},
	}
	byteData := uplinkMsg.FinalData
	var msgType manhole_cover.ManholeCoverMsgType
	// 心跳包小于13字节, 没有有用数据, 无需解析; 数据包大于13字节, 以此作为区分依据
	if len(byteData) > 13 {
		msgType = manhole_cover.ManholeCoverMsgType_Data
		parser00(byteData, uplinkMsgPayload)
	}

	// 解析完成后入库
	baseAppModel := &model.App{}
	err = parser.FillBaseAppModel(device.db, uplinkMsg, uplinkMsgPayload, baseAppModel)
	if err != nil {
		// 错误日志在FillBaseAppModel函数内部打印过了
		return
	}
	manholeCoverModel := &manhole_cover.ManholeCover{
		App:     *baseAppModel, // 引用类型
		MsgType: msgType,
	}

	err = device.db.Create(manholeCoverModel).Error
	if err != nil {
		device.logger.Errorf("create geomagnetic app message record failed in Decode(): %s", err.Error())
		return
	}
	// 如果井盖或水位异常, 创建异常通知
	err = parser.ReportManholeCoverNotification(device.db, manholeCoverModel)
	if err != nil {
		device.logger.Errorf("try to report manhole cover notification failed in Decode(): %s", err.Error())
		return
	}
	return
}

func fillModel(manholeCoverModel *manhole_cover.ManholeCover, uplinkMsgPayload *model.UplinkMsgPayload) (err error) {
	status, exist := uplinkMsgPayload.Attributes["Status"]
	if exist && status != nil {
		manholeCoverModel.Status = int(status.Value)
	}

	voltage, exist := uplinkMsgPayload.Attributes["Voltage"]
	if exist && voltage != nil {
		manholeCoverModel.Voltage = voltage.Value
	}

	return
}
