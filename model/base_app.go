package model

import "database/sql/driver"

// UplinkMsgPackageType 上行信息消息体类型, 可选值如下枚举示例
type UplinkMsgPackageType int

const (
	// PackageTypeStatus 有效状态消息
	PackageTypeStatus = iota
	// PackageTypeHeartbeat 通用心跳消息, 基本只能表示
	PackageTypeHeartbeat
	// PackageTypeConfig 上报当前设备的配置信息
	PackageTypeConfig
	// PackageTypeTriggered 事件触发, 导致状态变化上报
	// 其实可等同于Status类型, 这里细分一下. 唯传地磁设备有作区分
	PackageTypeTriggered
	// PackageTypeError 设备错误信息
	PackageTypeError
)

// UplinkMsgPayloadField 上行消息字段信息通用类型
type UplinkMsgPayloadField struct {
	DisplayName string  // 字段中文名: 温度, 温度, 压强等
	Value       float64 // 字段值: 37度(没有单位)
	Unit        string  // 单位: 摄氏度, 千帕等
}

// UplinkMsgAttribute 上行信息属性字段解析后的map映射集合
type UplinkMsgAttribute map[string]*UplinkMsgPayloadField

// UplinkMsgPayload 包含上行消息包类型, 及解析后的字段信息
type UplinkMsgPayload struct {
	Attributes UplinkMsgAttribute
}

// Scan 解析前操作
func (payload *UplinkMsgPayload) Scan(data interface{}) (err error) {
	return scan(data, payload)
}

// Value 存储前操作
func (payload *UplinkMsgPayload) Value() (driver.Value, error) {
	return value(payload)
}

// App 应用表基础模型
type App struct {
	// DeviceID非外键, 单纯表示应用表中单行记录的设备id
	DeviceSN   string `gorm:"not null"`
	Group      *Group
	GroupID    uint64 `gorm:"not null"`
	CustomerID uint64 `gorm:"not null"`

	FPort   uint8
	RawData string            // base64解密后的字节数组
	Data    *UplinkMsgPayload `gorm:"type:text"` // 结构解析后的数据字符串
}
