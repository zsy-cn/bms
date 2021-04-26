package geomagnetic

import (
	"github.com/jinzhu/gorm"
	"github.com/zsy-cn/bms/model"
)

var Tables = []interface{}{
	&Geomagnetic{},
}

func Migrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(Tables...).Error
	if err != nil {
		panic(err)
	}

	return
}

type GeomagneticMsgType int32

const (
	// GeomagneticMsgType_Type_Status 有效状态消息
	GeomagneticMsgType_Status GeomagneticMsgType = 0
	// GeomagneticMsgType_Heartbeat 通用心跳消息, 基本只能表示
	GeomagneticMsgType_Heartbeat GeomagneticMsgType = 1
	// GeomagneticMsgType_Heartbeat_New 新版心跳消息
	GeomagneticMsgType_Heartbeat_New GeomagneticMsgType = 2
	// GeomagneticMsgType_Config 上报当前设备的配置信息
	GeomagneticMsgType_Config GeomagneticMsgType = 3
	// GeomagneticMsgType_Triggered 事件触发, 导致状态变化上报
	// 其实可等同于Status类型, 这里细分一下. 唯传地磁设备有作区分
	GeomagneticMsgType_Triggered GeomagneticMsgType = 4
	// GeomagneticMsgType_Triggered_New 新版触发信息
	GeomagneticMsgType_Triggered_New GeomagneticMsgType = 5
	// GeomagneticMsgType_Error 设备错误信息
	GeomagneticMsgType_Error GeomagneticMsgType = 6
	// GeomagneticMsgType_Confirm confirm回复包
	GeomagneticMsgType_Confirm GeomagneticMsgType = 7
	// GeomagneticMsgType_XYZ 三轴数据上报帧
	GeomagneticMsgType_XYZ GeomagneticMsgType = 8
)

// 地磁当前状态表
type Geomagnetic struct {
	model.Base
	model.App
	MsgType GeomagneticMsgType
	XAxis   float64
	YAxis   float64
	ZAxis   float64
	Battery int     // 电量, 百分比, 整型
	Voltage float64 // 电压, 伏特, 浮点型
	Value   string
}
