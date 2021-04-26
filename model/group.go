package model

// Group 分组表(所有设备的分组, 包括各种传感器, IP音箱, 摄像头, 网关, 路由器)
type Group struct {
	Base
	Name         string `json:"name" gorm:"size:150"`
	DeviceTypeID uint64 `json:"deviceTypeId"`
	CustomerID   uint64 `json:"userId"` // 该分组属于哪个用户
	Status       uint8  `json:"status"` // [备用]状态标志位, 例如音箱设备中的播放分组, 0和1就够了.
}
