package model

import (
	"time"
)

// Base 基础模型
type Base struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt" gorm:"DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	// DeletedAt *time.Time `json:"deletedAt"`
}
