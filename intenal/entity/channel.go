package entity

import "time"

type PaymentChannel struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	PaymentMethodID *uint   `gorm:"index" json:"payment_method_id"`
	Code            string  `gorm:"type:varchar(255);unique;not null" json:"code"`
	Name            string  `gorm:"type:varchar(50);unique;not null" json:"name"`
	IconUrl         string  `gorm:"type:varchar(255)" json:"icon_url"`
	OrderNum        int     `gorm:"default:1" json:"order_num"`
	LibName         string  `gorm:"type:varchar(255)" json:"lib_name"`
	UserAction      string  `gorm:"type:varchar(25);not null" json:"user_action"`
	Mdr             string  `gorm:"type:varchar(255);default:'0'" json:"mdr"`
	FixedFee        float64 `gorm:"type:numeric;default:0" json:"fixed_fee"`

	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	PaymentMethod *PaymentMethod `gorm:"foreignKey:PaymentMethodID;references:ID" json:"payment_method"`
}
