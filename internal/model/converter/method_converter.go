package converter

import (
	"github.com/dwincahya/payment-be/internal/entity"
	models "github.com/dwincahya/payment-be/internal/model"
)

func PaymentMethodtoResponse(PaymentMethod *entity.PaymentMethod) *models.PaymentMethodResponse {
	return &models.PaymentMethodResponse{
		ID:         PaymentMethod.ID,
		Name:       PaymentMethod.Name,
		Desc:       PaymentMethod.Desc,
		Code:       PaymentMethod.Code,
		OrderNum:   PaymentMethod.OrderNum,
		UserAction: PaymentMethod.UserAction,
		CreatedAt:  PaymentMethod.CreatedAt,
		UpdatedAt:  PaymentMethod.UpdatedAt,
	}
}
