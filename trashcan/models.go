package trashcan

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/model"
)

var Tables = []interface{}{
	&Trashcan{},
}

func Migrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(Tables...).Error
	if err != nil {
		panic(err)
	}

	return
}

type TrashcanMsgType int32

const (
	// TrashcanMsgType_Register 注册信息
	TrashcanMsgType_Register TrashcanMsgType = 0
	// TrashcanMsgType_Status 状态信息(空/满, 距底部距离, 电量等)
	TrashcanMsgType_Status TrashcanMsgType = 1
	// TrashcanMsgType_Heartbeat 心跳信息(上盖角度, 距离, 电量等)
	TrashcanMsgType_Heartbeat TrashcanMsgType = 2
	// TrashcanMsgType_Position 位置信息(GPS, 电量等)
	TrashcanMsgType_Position TrashcanMsgType = 3
	// TrashcanMsgType_Status2 上盖状态信息(上盖是否打开, 打开角度, 电量等)
	TrashcanMsgType_Status2 TrashcanMsgType = 4
	// TrashcanMsgType_Reply 配置回复信息(应用的配置命令及参数)
	TrashcanMsgType_Reply TrashcanMsgType = 5
)

// 垃圾箱当前状态表
type Trashcan struct {
	model.Base
	model.App

	MsgType TrashcanMsgType
	Percent float64
	Status  int // [0: 掉线, 1: 正常, 2: 报警], 入库时根据规则判断
}
