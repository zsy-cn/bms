package model

// Manufacturer 传感器厂商表(预定义)
type Manufacturer struct {
	Base
	Name string `json:"name" gorm:"unique;not null;size:50"`
}
