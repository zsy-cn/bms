package model

// Rule 告警规则表
type Rule struct {
	Base
	CustomerID   uint64
	DeviceTypeID uint64
	DeviceSN     string        `gorm:"not null;unique"`
	Section      *Float64Array `gorm:"type:text"`
}
