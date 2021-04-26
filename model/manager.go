package model

// Manager ...
type Manager struct {
	Base
	Name        string `gorm:"not null;unique"`
	Passwd      string `gorm:"not null"`
	DisplayName string `gorm:"not null"`
	Phone       string `gorm:"unique"`
	RoleID      uint64 `gorm:"unique"`
	RoleEnable  bool   `gorm:"default:true"`
	Enable      bool   `gorm:"default:true"`
}
