package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/dwincahya/payment-be/internal/entity"
	models "github.com/dwincahya/payment-be/internal/model"
	"github.com/dwincahya/payment-be/internal/model/converter"
	"github.com/dwincahya/payment-be/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PaymentMethodUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	PaymentMethodRespository *repository.PaymentMethodRepository
}

func NewPaymentMethodUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, PaymentMethodRespository *repository.PaymentMethodRepository) *PaymentMethodUseCase {
	return &PaymentMethodUseCase{
		DB:                       db,
		Log:                      logger,
		Validate:                 validate,
		PaymentMethodRespository: PaymentMethodRespository,
	}
}

func (c *PaymentMethodUseCase) Create(ctx context.Context, request *models.CreatePaymentMethodRequest) (*models.PaymentMethodResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("Validation error")
		return nil, err
	}

	existingMethod := new(entity.PaymentMethod)
	if err := c.PaymentMethodRespository.FindByCode(tx, existingMethod, request.Code); err == nil {
		c.Log.Warnf("Payment method with code %s already exists", request.Code)
		return nil, errors.New("payment method with this code already exists")
	} else if err != gorm.ErrRecordNotFound {
		c.Log.WithError(err).Error("Failed to check existing payment method")
		return nil, err
	}

	paymentMethod := &entity.PaymentMethod{
		Name:       request.Name,
		Code:       request.Code,
		Desc:       request.Desc,
		OrderNum:   request.OrderNum,
		UserAction: request.UserAction,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := c.PaymentMethodRespository.Create(tx, paymentMethod); err != nil {
		c.Log.WithError(err).Error("Failed to create payment method")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, err
	}

	return converter.PaymentMethodtoResponse(paymentMethod), nil
}

func (c *PaymentMethodUseCase) Update(ctx context.Context, request *models.UpdatePaymentMethodRequest) (*models.PaymentMethodResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("Validation error")
		return nil, err
	}

	paymentMethod := new(entity.PaymentMethod)
	if err := c.PaymentMethodRespository.FindById(tx, paymentMethod, request.ID); err != nil {
		c.Log.WithError(err).Error("Payment method not found")
		return nil, err
	}

	paymentMethod.Name = request.Name
	paymentMethod.Code = request.Code
	paymentMethod.Desc = request.Desc
	paymentMethod.OrderNum = request.OrderNum
	paymentMethod.UserAction = request.UserAction
	paymentMethod.UpdatedAt = time.Now()

	if err := c.PaymentMethodRespository.Update(tx, paymentMethod); err != nil {
		c.Log.WithError(err).Error("Failed to update payment method")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, err
	}

	return converter.PaymentMethodtoResponse(paymentMethod), nil
}

func (c *PaymentMethodUseCase) Get(ctx context.Context, request *models.GetPaymentMethodRequest) (*models.PaymentMethodResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	paymentMethod := new(entity.PaymentMethod)
	if err := c.PaymentMethodRespository.FindById(tx, paymentMethod, request.ID); err != nil {
		c.Log.WithError(err).Error("Payment method not found")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, err
	}

	return converter.PaymentMethodtoResponse(paymentMethod), nil

}

func (c *PaymentMethodUseCase) Delete(ctx context.Context, request *models.DeletePaymentMethodRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	paymentMethod := new(entity.PaymentMethod)

	if err := c.PaymentMethodRespository.FindById(tx, paymentMethod, request.ID); err != nil {
		c.Log.WithError(err).Error("Payment method not found")
		return err
	}

	if err := c.PaymentMethodRespository.Delete(tx, paymentMethod); err != nil {
		c.Log.WithError(err).Error("Failed to delete payment method")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return err
	}
	return nil
}

func (c *PaymentMethodUseCase) List(ctx context.Context) ([]*models.PaymentMethodResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	paymentMethod, err := c.PaymentMethodRespository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("Failed to find all payment methods")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, err
	}

	response := make([]*models.PaymentMethodResponse, len(paymentMethod))
	for i, paymentMethod := range paymentMethod {
		response[i] = converter.PaymentMethodtoResponse(&paymentMethod)
	}

	return response, nil
}
