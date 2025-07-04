package models

import "time"

type PaymentMethodResponse struct {
	ID         uint      `json:"id"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	Desc       string    `json:"desc"`
	OrderNum   int       `json:"order_num"`
	UserAction string    `json:"user_action"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreatePaymentMethodRequest struct {
	Code       string `json:"code" validate:"required,max=25"`
	Name       string `json:"name" validate:"required,max=50"`
	Desc       string `json:"desc" validate:"required"`
	OrderNum   int    `json:"order_num" validate:"required"`
	UserAction string `json:"user_action" validate:"required,max=25"`
}

type GetPaymentMethodRequest struct {
	ID uint `json:"id" validate:"required"`
}

type UpdatePaymentMethodRequest struct {
	ID         uint   `json:"id" validate:"required"`
	Code       string `json:"code" validate:"required,max=25"`
	Name       string `json:"name" validate:"required,max=50"`
	Desc       string `json:"desc" validate:"required"`
	OrderNum   int    `json:"order_num" validate:"required"`
	UserAction string `json:"user_action" validate:"required,max=25"`
}

type DeletePaymentMethodRequest struct {
	ID uint `json:"id" validate:"required"`
}
