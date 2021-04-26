package model

// DeviceModel 设备型号表(预定义)
type DeviceModel struct {
	Base
	Name           string `json:"name" gorm:"unique;not null"`
	Description    string `json:"description"`
	ManufacturerID uint64 `json:"manufacturerId"`
	DeviceTypeID   uint64 `json:"deviceTypeId"`
}
