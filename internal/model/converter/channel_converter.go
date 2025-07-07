package converter

import (
	"github.com/dwincahya/payment-be/internal/entity"
	models "github.com/dwincahya/payment-be/internal/model"
)

func PaymentChanneltoResponse(PaymentChannel *entity.PaymentChannel) *models.PaymentChannelResponse {
	if PaymentChannel == nil {
		return nil
	}

	response := &models.PaymentChannelResponse{
		ID:              PaymentChannel.ID,
		Name:            PaymentChannel.Name,
		Code:            PaymentChannel.Code,
		PaymentMethodID: PaymentChannel.PaymentMethodID,
		OrderNum:        PaymentChannel.OrderNum,
		IconUrl:         PaymentChannel.IconUrl,
		LibName:         PaymentChannel.LibName,
		UserAction:      PaymentChannel.UserAction,
		Mdr:             PaymentChannel.Mdr,
		FixedFee:        PaymentChannel.FixedFee,
		PaymentMethod:   PaymentMethodtoResponse(PaymentChannel.PaymentMethod),
	}

	return response
}

func PaymentChanneltoResponseSlice(pcs []entity.PaymentChannel) []*models.PaymentChannelResponse {
	responses := make([]*models.PaymentChannelResponse, len(pcs))
	for i, pc := range pcs {
		responses[i] = PaymentChanneltoResponse(&pc)
	}
	return responses
}
