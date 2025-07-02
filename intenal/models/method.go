package models

type PaymentMethodResponse struct {
	ID         uint   `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	OrderNum   int64  `json:"order_num"`
	UserAction string `json:"user_action"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
type CreatePaymentMethodRequest struct {
	ID         uint   `json:"id" validate:"required"`
	Code       string `json:"code", validate:"required, max=255"`
	Name       string `json:"name" validate:"required,max=255"`
	Desc       string `json:"desc" validate:"required,max=255"`
	OrderNum   int64  `json:"order_num" validate:"required"`
	UserAction string `json:"user_action" validate:"required, max=255"`
	CreatedAt  string `json:"created_at" validate:"required"`
	UpdatedAt  string `json:"updated_at" validate:"required"`
}
