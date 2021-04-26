package model

// DeviceType 设备类型表(预定义)
// [地磁, 超声波垃圾箱, 环境监测, SOS, IP音箱, 摄像头, 路由器]
type DeviceType struct {
	Base
	Key      string `json:"key" gorm:"size:50;unique;not null"`
	Name     string `json:"name" gorm:"size:50;not null;"`
	IsSensor bool   `json:"isSensor"`
	// 属性表, 即这种设备需要哪些维度的信息, 只有传感器应用才需要
	Properties *StringArray `json:"properties" gorm:"type:text"`
}
