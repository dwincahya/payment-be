package repository

import (
	"github.com/dwincahya/payment-be/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PaymentMethodRepository struct {
	Repository[entity.PaymentMethod]
	Log *logrus.Logger
}

func NewPaymentMethodRepository(log *logrus.Logger) *PaymentMethodRepository {
	return &PaymentMethodRepository{
		Log: log,
	}
}

func (r *PaymentMethodRepository) FindByCode(tx *gorm.DB, paymentMethod *entity.PaymentMethod, code string) error {
	return tx.Where("code = ?", code).First(paymentMethod).Error
}

func (r *PaymentMethodRepository) FindById(tx *gorm.DB, paymentmethod *entity.PaymentMethod, id uint) error {
	return tx.Where("id = ?", id).First(paymentmethod).Error
}

func (r *PaymentMethodRepository) FindAll(tx *gorm.DB) ([]entity.PaymentMethod, error) {
	var paymentmethods []entity.PaymentMethod
	if err := tx.Find(&paymentmethods).Error; err != nil {
		r.Log.Error("Failed to find all payment methods: ", err)
		return nil, err
	}
	return paymentmethods, nil
}
