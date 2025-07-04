package usecase

import (
	"context"

	"github.com/dwincahya/payment-be/internal/entity"
	models "github.com/dwincahya/payment-be/internal/model"
	"github.com/dwincahya/payment-be/internal/model/converter"
	"github.com/dwincahya/payment-be/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PaymentChannelUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	PaymentChannelRepository *repository.PaymentChannelRepository
	PaymentMethodRepository  *repository.PaymentMethodRepository
}

func NewPaymentChannelUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, PaymentChannelRespository *repository.PaymentChannelRepository, PaymentMethodRespository *repository.PaymentMethodRepository) *PaymentChannelUseCase {
	return &PaymentChannelUseCase{
		DB:                       db,
		Log:                      logger,
		Validate:                 validate,
		PaymentChannelRepository: PaymentChannelRespository,
		PaymentMethodRepository:  PaymentMethodRespository,
	}
}

func (c *PaymentChannelUseCase) Create(ctx context.Context, request *models.CreatePaymentChannelRequest) (*models.PaymentChannelResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("Validation error for CreatePaymentChannelRequest")
		return nil, fiber.ErrBadRequest
	}

	paymentMethod := new(entity.PaymentMethod)

	if err := c.PaymentMethodRepository.FindById(tx, paymentMethod, *request.PaymentMethodID); err != nil {
		c.Log.WithError(err).Error("Parent Payment Method not found for PaymentChannel")
		return nil, fiber.ErrNotFound
	}

	existingChannel := new(entity.PaymentChannel)

	if err := c.PaymentChannelRepository.FindByCode(tx, existingChannel, request.Code); err == nil {
		c.Log.Warnf("Payment channel with code %s already exists", request.Code)
		return nil, fiber.NewError(fiber.StatusConflict, "Payment channel with this code already exists")
	} else if err != gorm.ErrRecordNotFound {
		c.Log.WithError(err).Error("Failed to check existing payment channel by code")
		return nil, fiber.ErrInternalServerError
	}

	paymentChannel := &entity.PaymentChannel{
		PaymentMethodID: request.PaymentMethodID,
		Code:            request.Code,
		Name:            request.Name,
		IconUrl:         request.IconUrl,
		OrderNum:        request.OrderNum,
		LibName:         request.LibName,
		UserAction:      request.UserAction,
		Mdr:             request.Mdr,
		FixedFee:        request.FixedFee,
	}

	if err := c.PaymentChannelRepository.Create(tx, paymentChannel); err != nil {
		c.Log.WithError(err).Error("Failed to create payment channel")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction for creating payment channel")
		return nil, fiber.ErrInternalServerError
	}

	return converter.PaymentChanneltoResponse(paymentChannel), nil

}

func (c *PaymentChannelUseCase) Update(ctx context.Context, request *models.UpdatePaymentChannelRequest) (*models.PaymentChannelResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("Validation error for UpdatePaymentChannelRequest")
		return nil, fiber.ErrBadRequest
	}

	paymentChannel := new(entity.PaymentChannel)
	if err := c.PaymentChannelRepository.FindById(tx, paymentChannel, request.ID); err != nil {
		c.Log.WithError(err).Error("Payment channel not found for update")
		return nil, fiber.ErrNotFound
	}

	if request.PaymentMethodID != nil && *request.PaymentMethodID != *paymentChannel.PaymentMethodID {
		paymentMethod := new(entity.PaymentMethod)
		if err := c.PaymentMethodRepository.FindById(tx, paymentMethod, *request.PaymentMethodID); err != nil {
			c.Log.WithError(err).Error("New Payment Method not found for PaymentChannel update")
			return nil, fiber.ErrNotFound
		}
	}

	paymentChannel.PaymentMethodID = request.PaymentMethodID
	paymentChannel.Code = request.Code
	paymentChannel.Name = request.Name
	paymentChannel.IconUrl = request.IconUrl
	paymentChannel.OrderNum = request.OrderNum
	paymentChannel.LibName = request.LibName
	paymentChannel.UserAction = request.UserAction
	paymentChannel.Mdr = request.Mdr
	paymentChannel.FixedFee = request.FixedFee

	if err := c.PaymentChannelRepository.Update(tx, paymentChannel); err != nil {
		c.Log.WithError(err).Error("Failed to update payment channel")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.PaymentChanneltoResponse(paymentChannel), nil
}

func (c *PaymentChannelUseCase) Get(ctx context.Context, request *models.GetPaymentChannelRequest) (*models.PaymentChannelResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	paymentChannel := new(entity.PaymentChannel)

	if err := c.PaymentChannelRepository.FindById(tx.Preload("PaymentMethod"), paymentChannel, request.ID); err != nil {
		c.Log.WithError(err).Error("Payment channel not found")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction for getting payment channel")
		return nil, fiber.ErrInternalServerError
	}

	return converter.PaymentChanneltoResponse(paymentChannel), nil
}

func (c *PaymentChannelUseCase) Delete(ctx context.Context, request *models.DeletePaymentChannelRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	paymentChannel := new(entity.PaymentChannel)
	if err := c.PaymentChannelRepository.FindById(tx, paymentChannel, request.ID); err != nil {
		c.Log.WithError(err).Error("Payment channel not found for deletion")
		return fiber.ErrNotFound
	}

	if err := c.PaymentChannelRepository.Delete(tx, paymentChannel); err != nil {
		c.Log.WithError(err).Error("Failed to delete payment channel")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction for deleting payment channel")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *PaymentChannelUseCase) List(ctx context.Context, request *models.ListPaymentChannelRequest) ([]*models.PaymentChannelResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	query := tx.Preload("PaymentMethod")
	if request.PaymentMethodID != nil {
		query = query.Where("payment_method_id = ?", *request.PaymentMethodID)
	}

	if request.Page > 0 && request.Limit > 0 {
		offset := (request.Page - 1) * request.Limit
		query = query.Offset(offset).Limit(request.Limit)
	}

	paymentChannels, err := c.PaymentChannelRepository.FindAll(query)
	if err != nil {
		c.Log.WithError(err).Error("Failed to find all payment channels")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.PaymentChanneltoResponseSlice(paymentChannels), nil
}
