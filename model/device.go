package model

// Device 设备通用属性(各种传感器, IP音箱, 摄像头, 网关, 路由器)
type Device struct {
	Base
	Name         string `json:"name" gorm:"size:50;not null"`
	SerialNumber string `json:"serialNumber" gorm:"size:100;unique"`
	Description  string `json:"description" gorm:"size:100"`
	Position     string `gorm:"size:100;"`

	Customer      *Customer    `json:"customer"`
	CustomerID    uint64       `json:"customerId" gorm:"not null"`
	Group         *Group       `json:"group"`
	GroupID       uint64       `json:"groupId" gorm:"not null"`
	DeviceType    *DeviceType  `json:"deviceType"`
	DeviceTypeID  uint64       `json:"deviceTypeId" gorm:"not null"`
	DeviceModel   *DeviceModel `json:"deviceModel"`
	DeviceModelID uint64       `json:"deviceModelId" gorm:"not null"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Actived   bool    `gorm:"default:true"`
}
