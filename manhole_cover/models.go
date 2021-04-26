package manhole_cover

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/model"
)

var Tables = []interface{}{
	&ManholeCover{},
}

func Migrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(Tables...).Error
	if err != nil {
		panic(err)
	}

	return
}

type ManholeCoverMsgType int32

const (
	// ManholeCoverMsgType_Heartbeat 心跳包
	ManholeCoverMsgType_Heartbeat ManholeCoverMsgType = 0
	// ManholeCoverMsgType_Data 数据包
	ManholeCoverMsgType_Data ManholeCoverMsgType = 1
)

// 井盖当前状态表
type ManholeCover struct {
	model.Base
	model.App
	MsgType          ManholeCoverMsgType
	Status           int `json:"status"`
	WaterLevelStatus int // 水位报警
	Voltage          float64
}
