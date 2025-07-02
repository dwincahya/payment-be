package models

type PaymentMethodResponse struct {
	ID         uint   `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	OrderNum   int64  `json:"order_num"`
	UserAction string `json:"user_action"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}
type CreatePaymentMethodRequest struct {
	ID       uint   `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required,max=255"`
	Desc     string `json:"desc" validate:"required,max=255"`
	OrderNum int64  `json:"order_num" validate:"required"`
	User
}
