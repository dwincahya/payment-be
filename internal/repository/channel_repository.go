package repository

import (
	"github.com/dwincahya/payment-be/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PaymentChannelRepository struct {
	Repository[entity.PaymentChannel]
	Log *logrus.Logger
}

func NewPaymentChannelRepository(log *logrus.Logger) *PaymentChannelRepository {
	return &PaymentChannelRepository{
		Log: log,
	}
}

func (r *PaymentChannelRepository) FindById(tx *gorm.DB, paymentChannel *entity.PaymentChannel, id uint) error {
	return tx.Where("id = ?", id).First(paymentChannel).Error
}

func (r *PaymentChannelRepository) FindAllByPaymentMethodId(tx *gorm.DB, paymentMethodId uint) ([]entity.PaymentChannel, error) {
	var paymentChannels []entity.PaymentChannel
	if err := tx.Where("payment_method_id = ?", paymentMethodId).Find(&paymentChannels).Error; err != nil {
		r.Log.Error("Failed to find payment channels by payment method ID: ", err)
		return nil, err
	}
	return paymentChannels, nil
}

func (r *PaymentChannelRepository) FindByCode(tx *gorm.DB, paymentChannel *entity.PaymentChannel, code string) error {
	return tx.Where("code = ?", code).First(paymentChannel).Error
}

func (r *PaymentChannelRepository) FindAll(tx *gorm.DB) ([]entity.PaymentChannel, error) {
	var paymentChannels []entity.PaymentChannel
	if err := tx.Find(&paymentChannels).Error; err != nil {
		r.Log.Error("Failed to find all payment Channels: ", err)
		return nil, err
	}
	return paymentChannels, nil
}
