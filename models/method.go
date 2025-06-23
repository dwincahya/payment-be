package models

import (
	"time"
)

type PaymentMethod struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Name       string     `gorm:"type:varchar(50);unique;not null" json:"name"`
	Desc       string     `gorm:"type:text" json:"desc"`
	OrderNum   int        `gorm:"default:1;not null" json:"order_num"`
	UserAction string     `gorm:"type:varchar(25);not null" json:"user_action"`
	Code       string     `gorm:"type:varchar(25)" json:"code"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`

	Channels []PaymentChannel `gorm:"foreignKey:PaymentMethodID" json:"channels,omitempty"`
}
