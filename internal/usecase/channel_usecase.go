package usecase

import (
	"context"
	"errors"
	"strconv"

	"github.com/dwincahya/payment-be/internal/entity"
	models "github.com/dwincahya/payment-be/internal/model"
	"github.com/dwincahya/payment-be/internal/model/converter"
	"github.com/dwincahya/payment-be/internal/repository"
	"github.com/go-playground/validator/v10"
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
		return nil, err
	}

	paymentMethod := new(entity.PaymentMethod)

	if err := c.PaymentMethodRepository.FindById(tx, paymentMethod, *request.PaymentMethodID); err != nil {
		c.Log.WithError(err).Error("Parent Payment Method not found for PaymentChannel")
		return nil, err
	}

	existingChannel := new(entity.PaymentChannel)

	if err := c.PaymentChannelRepository.FindByCode(tx, existingChannel, request.Code); err == nil {
		c.Log.Warnf("Payment channel with code %s already exists", request.Code)
		return nil, errors.New("payment channel with this code already exists")
	} else if err != gorm.ErrRecordNotFound {
		c.Log.WithError(err).Error("Failed to check existing payment channel by code")
		return nil, errors.New("failed to check existing payment channel")
	}

	paymentChannel := &entity.PaymentChannel{
		PaymentMethodID: request.PaymentMethodID,
		Code:            request.Code,
		Name:            request.Name,
		IconUrl:         request.IconUrl,
		OrderNum:        request.OrderNum,
		LibName:         request.LibName,
		UserAction:      request.UserAction,
		Mdr:             strconv.Itoa(request.Mdr),
		FixedFee:        request.FixedFee,
	}

	if err := c.PaymentChannelRepository.Create(tx, paymentChannel); err != nil {
		c.Log.WithError(err).Error("Failed to create payment channel")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction for creating payment channel")
		return nil, err
	}

	return converter.PaymentChanneltoResponse(paymentChannel), nil

}

func (c *PaymentChannelUseCase) Update(ctx context.Context, request *models.UpdatePaymentChannelRequest) (*models.PaymentChannelResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("Validation error for UpdatePaymentChannelRequest")
		return nil, err
	}

	paymentChannel := new(entity.PaymentChannel)
	if err := c.PaymentChannelRepository.FindById(tx, paymentChannel, request.ID); err != nil {
		c.Log.WithError(err).Error("Payment channel not found for update")
		return nil, err
	}

	if request.PaymentMethodID != nil && *request.PaymentMethodID != *paymentChannel.PaymentMethodID {
		paymentMethod := new(entity.PaymentMethod)
		if err := c.PaymentMethodRepository.FindById(tx, paymentMethod, *request.PaymentMethodID); err != nil {
			c.Log.WithError(err).Error("New Payment Method not found for PaymentChannel update")
			return nil, err
		}
	}

	paymentChannel.PaymentMethodID = request.PaymentMethodID
	paymentChannel.Code = request.Code
	paymentChannel.Name = request.Name
	paymentChannel.IconUrl = request.IconUrl
	paymentChannel.OrderNum = request.OrderNum
	paymentChannel.LibName = request.LibName
	paymentChannel.UserAction = request.UserAction
	paymentChannel.Mdr = strconv.Itoa(request.Mdr)
	paymentChannel.FixedFee = request.FixedFee

	if err := c.PaymentChannelRepository.Update(tx, paymentChannel); err != nil {
		c.Log.WithError(err).Error("Failed to update payment channel")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, err
	}

	return converter.PaymentChanneltoResponse(paymentChannel), nil
}

func (c *PaymentChannelUseCase) Get(ctx context.Context, request *models.GetPaymentChannelRequest) (*models.PaymentChannelResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	paymentChannel := new(entity.PaymentChannel)

	if err := c.PaymentChannelRepository.FindById(tx.Preload("PaymentMethod"), paymentChannel, request.ID); err != nil {
		c.Log.WithError(err).Error("Payment channel not found")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction for getting payment channel")
		return nil, err
	}

	return converter.PaymentChanneltoResponse(paymentChannel), nil
}

func (c *PaymentChannelUseCase) Delete(ctx context.Context, request *models.DeletePaymentChannelRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	paymentChannel := new(entity.PaymentChannel)
	if err := c.PaymentChannelRepository.FindById(tx, paymentChannel, request.ID); err != nil {
		c.Log.WithError(err).Error("Payment channel not found for deletion")
		return err
	}

	if err := c.PaymentChannelRepository.Delete(tx, paymentChannel); err != nil {
		c.Log.WithError(err).Error("Failed to delete payment channel")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction for deleting payment channel")
		return err
	}

	return nil
}

func (c *PaymentChannelUseCase) List(ctx context.Context, request *models.ListPaymentChannelRequest) ([]*models.PaymentChannelResponse, *models.PageMetadata, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	query := tx.Model(&entity.PaymentChannel{})

	if request.PaymentMethodID != nil {
		query = query.Where("payment_method_id = ?", *request.PaymentMethodID)
	}

	if request.Code != "" {
		query = query.Where("LOWER(code) LIKE LOWER(?)", "%"+request.Code+"%")
	}

	if request.Name != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+request.Name+"%")
	}

	var totalRows int64
	if err := query.Count(&totalRows).Error; err != nil {
		c.Log.WithError(err).Error("Failed to count payment channels")
		return nil, nil, err
	}

	paymentQuery := query.Preload("PaymentMethod")

	if !request.All {
		if request.Page > 0 && request.Limit > 0 {
			offset := (request.Page - 1) * request.Limit
			paymentQuery = paymentQuery.Offset(offset).Limit(request.Limit)
		}
	}

	var paymentChannels []entity.PaymentChannel
	if err := paymentQuery.Find(&paymentChannels).Error; err != nil {
		c.Log.WithError(err).Error("Failed to find payment channels")
		return nil, nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, nil, err
	}

	totalPages := 0
	if request.Limit > 0 && !request.All {
		totalPages = int((totalRows + int64(request.Limit) - 1) / int64(request.Limit))
	} else if request.All {
		totalPages = 1
	}

	paging := &models.PageMetadata{
		Page:      request.Page,
		Limit:     request.Limit,
		TotalItem: int(totalRows),
		TotalPage: totalPages,
	}

	responseData := converter.PaymentChanneltoResponseSlice(paymentChannels)

	return responseData, paging, nil
}
