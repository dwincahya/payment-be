package models

import "time"

type PaymentChannelResponse struct {
	ID              uint                   `json:"id"`
	PaymentMethodID *uint                  `json:"payment_method_id"`
	Code            string                 `json:"code"`
	Name            string                 `json:"name"`
	IconUrl         string                 `json:"icon_url"`
	OrderNum        int                    `json:"order_num"`
	LibName         string                 `json:"lib_name"`
	UserAction      string                 `json:"user_action"`
	Mdr             string                 `json:"mdr"`
	FixedFee        float64                `json:"fixed_fee"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
	PaymentMethod   *PaymentMethodResponse `json:"payment_method,omitempty"`
}

type CreatePaymentChannelRequest struct {
	PaymentMethodID *uint   `json:"payment_method_id" validate:"required"`
	Code            string  `json:"code" validate:"required,max=255"`
	Name            string  `json:"name" validate:"required,max=50"`
	IconUrl         string  `json:"icon_url" validate:"required,max=255"`
	OrderNum        int     `json:"order_num" validate:"required"`
	LibName         string  `json:"lib_name" validate:"required,max=255"`
	UserAction      string  `json:"user_action" validate:"required,max=25"`
	Mdr             string  `json:"mdr" validate:"required,max=255"`
	FixedFee        float64 `json:"fixed_fee" validate:"required"`
}

type GetPaymentChannelRequest struct {
	ID              uint  `json:"id" validate:"required"`
	PaymentMethodID *uint `json:"payment_method_id" params:"payment_method_id"`
}

type ListPaymentChannelRequest struct {
	ID              *uint `json:"id,omitempty" query:"id"`
	PaymentMethodID *uint `json:"payment_method_id,omitempty" query:"payment_method_id"`
	Page            int   `json:"page" query:"page"`
	Limit           int   `json:"size" query:"size"`
}

type UpdatePaymentChannelRequest struct {
	ID              uint    `json:"id" validate:"required"`
	PaymentMethodID *uint   `json:"payment_method_id" validate:"required"`
	Code            string  `json:"code" validate:"required,max=255"`
	Name            string  `json:"name" validate:"required,max=50"`
	IconUrl         string  `json:"icon_url" validate:"required,max=255"`
	OrderNum        int     `json:"order_num" validate:"required"`
	LibName         string  `json:"lib_name" validate:"required,max=255"`
	UserAction      string  `json:"user_action" validate:"required,max=25"`
	Mdr             string  `json:"mdr" validate:"required,max=255"`
	FixedFee        float64 `json:"fixed_fee" validate:"required"`
}

type DeletePaymentChannelRequest struct {
	ID uint `json:"id" validate:"required"`
}
