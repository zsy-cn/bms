package model

// Contact 联系人表, 用户相关
type Contact struct {
	Base
	CustomerID uint64 `json:"customerId" gorm:"not null"`
	Name       string `json:"name" gorm:"size:20;not null"`
	Phone      string `json:"phone" gorm:"size:20"`
	Email      string `json:"email" gorm:"size:50"`
}

// ContactJSON ...
type ContactJSON struct {
	ID         uint64    `json:"id"`
	Customer   *Customer `json:"customer"`
	CustomerID uint64    `json:"customerId"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
}
