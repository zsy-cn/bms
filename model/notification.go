package model

import "time"

// Notification 通知表, 状态表触发规则后统计进入这里
type Notification struct {
	Base
	DeviceSN     string `gorm:"not null"`
	MsgID        uint64
	GroupID      uint64
	CustomerID   uint64
	DeviceTypeID uint64    // 水位/垃圾箱
	RuleID       uint64    // 规则ID, 水位和垃圾箱需要有此项, 井盖, SOS和烟感则不需要
	Status       int8      // [1: 通知(初级), 2: 警告(中级), 3: 严重警告(高级)]
	Key          string    // 固定告警类型, 如"manhole_cover_open_alert", "manhole_cover_water_alert", 这种, 同一种设备拥有多种类型的固定告警, 区别于rules.
	Content      string    // 提示内容, 不同设备类型的内容不同
	Solved       bool      `gorm:"default:false"` // 是否已处理
	Times        int8      // 告警次数, 同一个设备的告警在短时间内重复不会创建多少记录(目的是不会重复提示)
	SolvedAt     time.Time // 处理时间
	// SolvedBy // 处理人
}
