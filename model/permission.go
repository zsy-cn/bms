package model

// Permission ...
type Permission struct {
	Base
	Name string `gorm:"not null"`
	Path string `gorm:"not null;unique"` // url路径
}
