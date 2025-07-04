package usecase

import (
	"context"
	"time"

	"github.com/dwincahya/payment-be/internal/entity"
	models "github.com/dwincahya/payment-be/internal/model"
	"github.com/dwincahya/payment-be/internal/model/converter"
	"github.com/dwincahya/payment-be/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
		return nil, fiber.ErrBadRequest
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
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.PaymentMethodtoResponse(paymentMethod), nil
}

func (c *PaymentMethodUseCase) Update(ctx context.Context, request *models.UpdatePaymentMethodRequest) (*models.PaymentMethodResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("Validation error")
		return nil, fiber.ErrBadRequest
	}

	paymentMethod := new(entity.PaymentMethod)
	if err := c.PaymentMethodRespository.FindById(tx, paymentMethod, request.ID); err != nil {
		c.Log.WithError(err).Error("Payment method not found")
		return nil, fiber.ErrNotFound
	}

	paymentMethod.Name = request.Name
	paymentMethod.Code = request.Code
	paymentMethod.Desc = request.Desc
	paymentMethod.OrderNum = request.OrderNum
	paymentMethod.UserAction = request.UserAction
	paymentMethod.UpdatedAt = time.Now()

	if err := c.PaymentMethodRespository.Update(tx, paymentMethod); err != nil {
		c.Log.WithError(err).Error("Failed to update payment method")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.PaymentMethodtoResponse(paymentMethod), nil
}

func (c *PaymentMethodUseCase) Get(ctx context.Context, request *models.GetPaymentMethodRequest) (*models.PaymentMethodResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	paymentMethod := new(entity.PaymentMethod)
	if err := c.PaymentMethodRespository.FindById(tx, paymentMethod, request.ID); err != nil {
		c.Log.WithError(err).Error("Payment method not found")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.PaymentMethodtoResponse(paymentMethod), nil

}

func (c *PaymentMethodUseCase) Delete(ctx context.Context, request *models.DeletePaymentMethodRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	paymentMethod := new(entity.PaymentMethod)

	if err := c.PaymentMethodRespository.FindById(tx, paymentMethod, request.ID); err != nil {
		c.Log.WithError(err).Error("Payment method not found")
		return fiber.ErrNotFound
	}

	if err := c.PaymentMethodRespository.Delete(tx, paymentMethod); err != nil {
		c.Log.WithError(err).Error("Failed to delete payment method")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return fiber.ErrInternalServerError
	}
	return nil
}

func (c *PaymentMethodUseCase) List(ctx context.Context) ([]*models.PaymentMethodResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	paymentMethod, err := c.PaymentMethodRespository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("Failed to find all payment methods")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	response := make([]*models.PaymentMethodResponse, len(paymentMethod))
	for i, paymentMethod := range paymentMethod {
		response[i] = converter.PaymentMethodtoResponse(&paymentMethod)
	}

	return response, nil
}
