package models

import (
	"time"
)

type PaymentMethod struct {
	ID         uint       `gorm:"primaryKey" json:"Id"`
	Name       string     `gorm:"type:varchar(50);unique;not null" json:"Name"`
	Desc       string     `gorm:"type:text" json:"Desc"`
	OrderNum   int        `gorm:"default:1;not null" json:"OrderNum"`
	UserAction string     `gorm:"type:varchar(25);not null" json:"UserAction"`
	Code       string     `gorm:"type:varchar(25)" json:"Code"`
	CreatedAt  *time.Time `json:"CreatedAt"`
	UpdatedAt  *time.Time `json:"UpdatedAt"`

	Channels []PaymentChannel `gorm:"foreignKey:PaymentMethodID" json:"channels,omitempty"`
}
