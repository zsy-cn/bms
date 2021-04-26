package sos_yingfeichi_01

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/conf"
	"github.com/zsy-cn/bms/model"
	"github.com/zsy-cn/bms/protos"
	"github.com/zsy-cn/bms/service/parser"
	"github.com/zsy-cn/bms/sos"
	"github.com/zsy-cn/bms/util/log"
)

// SOSYingfeichi01 ...
type SOSYingfeichi01 struct {
	logger *log.Logger
	db     *gorm.DB
}

// New ...
func New(logger *log.Logger) (device *SOSYingfeichi01, err error) {
	db, err := conf.ConnectDB()
	if err != nil {
		logger.Errorf("connect database failed: %s", err.Error())
		return
	}
	device = &SOSYingfeichi01{
		logger: logger,
		db:     db,
	}
	return
}

// Decode ...
// 要实现一个interface, receiver不能是指针???(SOSYingfeichi01实现api.go->Parser接口)
// uplinkMsg中FinalData已经经过base64解码, 无需重复操作, 直接解析即可
func (device SOSYingfeichi01) Decode(uplinkMsg *protos.ParserHubUplinkMsg) (err error) {
	device.logger.Debugf("receive uplink mesage in Decode(): %+v", uplinkMsg)
	uplinkMsgPayload := &model.UplinkMsgPayload{
		Attributes: model.UplinkMsgAttribute{},
	}
	// byteData := uplinkMsg.FinalData

	baseAppModel := &model.App{}
	err = parser.FillBaseAppModel(device.db, uplinkMsg, uplinkMsgPayload, baseAppModel)
	if err != nil {
		// 错误日志在FillBaseAppModel函数内部打印过了
		return
	}
	sosModel := &sos.Sos{
		App:    *baseAppModel, // 引用类型
		Status: 1,
	}

	err = device.db.Create(sosModel).Error
	if err != nil {
		device.logger.Errorf("create sos app message record failed in Decode(): %s", err.Error())
	}
	return
}

// HealthCheck ...
func (device SOSYingfeichi01) HealthCheck(empty *protos.Empty) (resp *protos.ParserHealthCheckResponse, err error) {
	resp = &protos.ParserHealthCheckResponse{
		Msg: "success",
	}
	return
}
