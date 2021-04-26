package model

// Role ...
type Role struct {
	Base
	Name          string       `gorm:"not null;unique"`
	Enable        bool         `gorm:"default:true"`
	PermissionIDs *Uint64Array `gorm:"type:text"`
}
