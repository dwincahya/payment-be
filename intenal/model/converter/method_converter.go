package converter

import (
	"github.com/dwincahya/payment-be/intenal/entity"
	models "github.com/dwincahya/payment-be/intenal/model"
)

func PaymentMethodtoResponse(PaymentMethod *entity.PaymentMethod) *models.PaymentMethodResponse {
	return &models.PaymentMethodResponse{
		ID:         PaymentMethod.ID,
		Name:       PaymentMethod.Name,
		Desc:       PaymentMethod.Desc,
		Code:       PaymentMethod.Code,
		OrderNum:   PaymentMethod.OrderNum,
		UserAction: PaymentMethod.UserAction,
	}
}
