package converter

import (
	"github.com/dwincahya/payment-be/internal/entity"
	models "github.com/dwincahya/payment-be/internal/model"
)

func PaymentMethodtoResponse(PaymentMethod *entity.PaymentMethod) *models.PaymentMethodResponse {

	if PaymentMethod == nil {
		return nil
	}
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

func PaymentMethodtoResponseSlice(entities []entity.PaymentMethod) []*models.PaymentMethodResponse {
	responses := make([]*models.PaymentMethodResponse, len(entities))
	for i, pm := range entities {
		responses[i] = PaymentMethodtoResponse(&pm)
	}
	return responses
}
