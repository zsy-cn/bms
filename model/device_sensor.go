package model

import "time"

// Sensor 传感器表
type Sensor struct {
	DeviceSN string `gorm:"size:100;unique"`
	DevEUI   string `gorm:"size:16;column:dev_eui"`
	AppEUI   string `gorm:"size:16;column:app_eui"`
	AppKey   string `gorm:"size:32"`
	Freq     string
}

// SensorJSON ...
type SensorJSON struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serialNumber"`
	Description  string `json:"description"`

	Group         string `json:"group"`
	GroupID       uint64 `json:"groupId"`
	DeviceType    string `json:"deviceType"`
	DeviceTypeID  uint64 `json:"deviceTypeId"`
	DeviceModel   string `json:"deviceModel"`
	DeviceModelID uint64 `json:"deviceModelId"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Status    uint8   `json:"status"` // [0: 断线, 1: 在线...其他状态, 由设备类型决定]

	DevEUI   string    `json:"devEUI"`
	AppEUI   string    `json:"appEUI"`
	AppKey   string    `json:"appkey"`
	Freq     string    `json:"freq"`
	LastSeen time.Time `json:"lastSeen"`
}
