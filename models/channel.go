package models

import (
	"time"
)

type PaymentChannel struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	PaymentMethodID *uint   `gorm:"index" json:"PaymentMethodID"`
	Code            string  `gorm:"type:varchar(255);unique;not null" json:"Code"`
	Name            string  `gorm:"type:varchar(50);unique;not null" json:"Name"`
	IconUrl         string  `gorm:"type:varchar(255)" json:"IconUrl"`
	OrderNum        int     `gorm:"default:1" json:"OrderNum"`
	LibName         string  `gorm:"type:varchar(255)" json:"LibName"`
	UserAction      string  `gorm:"type:varchar(25);not null" json:"UserAction"`
	Mdr             string  `gorm:"type:varchar(255);default:'0'" json:"Mdr"`
	FixedFee        float64 `gorm:"type:numeric;default:0" json:"FixedFee"`

	CreatedAt     *time.Time     `json:"CreatedAt"`
	UpdatedAt     *time.Time     `json:"UpdatedAt"`
	PaymentMethod *PaymentMethod `gorm:"foreignKey:PaymentMethodID;references:ID" json:"PaymentMethod"`

	PaymentMethodName  string `gorm:"-" json:"paymentMethod"`
	PaymentMethodValue uint   `gorm:"-" json:"paymentMethodValue"`
}
