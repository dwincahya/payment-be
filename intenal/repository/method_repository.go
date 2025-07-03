package repository

import (
	"fmt"
	"time"

	"github.com/dwincahya/payment-be/intenal/entity"
	"gorm.io/gorm"
)

type paymentMethodRepository struct {
	db *gorm.DB
}

type PaymentMethodRepository interface {
	Create(paymentMethod *entity.PaymentMethod) error
	FindByID(id uint) (*entity.PaymentMethod, error)
	FindAll() ([]entity.PaymentMethod, error)
	Update(paymentMethod *entity.PaymentMethod) error
	FindAllWithPagination(offset, limit int) ([]entity.PaymentMethod, int64, error)
	Delete(id uint) error
}

func NewPaymentMethodRepository(db *gorm.DB) *paymentMethodRepository {
	return &paymentMethodRepository{db: db}
}

func (r *paymentMethodRepository) Create(paymentMethod *entity.PaymentMethod) error {
	now := time.Now()
	if paymentMethod.CreatedAt == nil {
		paymentMethod.CreatedAt = &now
	}
	if paymentMethod.UpdatedAt == nil {
		paymentMethod.UpdatedAt = &now
	}

	return r.db.Create(paymentMethod).Error
}

func (r *paymentMethodRepository) FindByID(id uint) (*entity.PaymentMethod, error) {
	var pm entity.PaymentMethod
	if err := r.db.Preload("Channels").First(&pm, id).Error; err != nil {
		return nil, err
	}
	return &pm, nil
}

func (r *paymentMethodRepository) FindAll() ([]entity.PaymentMethod, error) {
	var pms []entity.PaymentMethod
	if err := r.db.Preload("Channels").Find(&pms).Error; err != nil {
		return nil, err
	}
	return pms, nil
}

func (r *paymentMethodRepository) Update(paymentMethod *entity.PaymentMethod) error {
	now := time.Now()
	paymentMethod.UpdatedAt = &now

	result := r.db.Save(paymentMethod)
	if result.Error != nil {
		return fmt.Errorf("failed to update payment method: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("payment method with ID %d not found for update", paymentMethod.ID)
	}
	return nil
}

func (r *paymentMethodRepository) FindAllWithPagination(offset, limit int) ([]entity.PaymentMethod, int64, error) {
	var pms []entity.PaymentMethod
	var total int64

	if err := r.db.Model(&entity.PaymentMethod{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count total payment methods: %w", err)
	}

	if err := r.db.Preload("Channels").Offset(offset).Limit(limit).Find(&pms).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to retrieve payment methods with pagination: %w", err)
	}

	return pms, total, nil
}

func (r *paymentMethodRepository) Delete(id uint) error {
	result := r.db.Delete(&entity.PaymentMethod{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete payment method: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("payment method with ID %d not found for deletion", id)
	}
	return nil
}
