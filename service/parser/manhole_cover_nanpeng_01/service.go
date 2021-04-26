package manhole_cover_nanpeng_01

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/manhole_cover"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/parser"
	"github.com/zsy-cn/bms/util/log"
)

// ManholeCoverNanpeng01 ...
type ManholeCoverNanpeng01 struct {
	parser.ConsulParserService
	logger *log.Logger
	db     *gorm.DB
}

// New ...
func New(logger *log.Logger) (device *ManholeCoverNanpeng01, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}
	device = &ManholeCoverNanpeng01{
		logger: logger,
		db:     db,
	}
	return
}

// Decode ...
// 要实现一个interface, receiver不能是指针???(ManholeCoverNanpeng01实现api.go->Parser接口)
// uplinkMsg中FinalData已经经过base64解码, 无需重复操作, 直接解析即可
func (device ManholeCoverNanpeng01) Decode(uplinkMsg *protos.ParserHubUplinkMsg) (err error) {
	device.logger.Debugf("receive uplink mesage in Decode(): %+v", uplinkMsg)
	device.logger.Debug("uplink mesage payload bytes in Decode(): ", uplinkMsg.FinalData)

	uplinkMsgPayload := &model.UplinkMsgPayload{
		Attributes: model.UplinkMsgAttribute{},
	}
	byteData := uplinkMsg.FinalData

	header := byteData[0]
	if header != 0x4E {
		device.logger.Error("unknown package, drop it")
		return
	}
	// 通用解析
	parser00(byteData, uplinkMsgPayload)

	// 解析完成后入库
	baseAppModel := &model.App{}
	err = parser.FillBaseAppModel(device.db, uplinkMsg, uplinkMsgPayload, baseAppModel)
	if err != nil {
		// 错误日志在FillBaseAppModel函数内部打印过了
		return
	}
	manholeCoverModel := &manhole_cover.ManholeCover{
		App:     *baseAppModel, // 引用类型
		MsgType: manhole_cover.ManholeCoverMsgType_Heartbeat,
	}

	err = fillModel(manholeCoverModel, uplinkMsgPayload)
	if err != nil {
		device.logger.Errorf("fill manhole_cover model failed in Decode(): %s", err.Error())
		return
	}

	err = device.db.Create(manholeCoverModel).Error
	if err != nil {
		device.logger.Errorf("create manhole_cover app message record failed in Decode(): %s", err.Error())
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

	waterLevelStatus, exist := uplinkMsgPayload.Attributes["WaterLevelStatus"]
	if exist && waterLevelStatus != nil {
		manholeCoverModel.WaterLevelStatus = int(waterLevelStatus.Value)
	}

	voltage, exist := uplinkMsgPayload.Attributes["Voltage"]
	if exist && voltage != nil {
		manholeCoverModel.Voltage = voltage.Value
	}

	return
}
