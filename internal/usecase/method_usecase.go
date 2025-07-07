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

func (c *PaymentMethodUseCase) List(ctx context.Context, request *models.ListPaymentMethodRequest) ([]*models.PaymentMethodResponse, *models.PageMetadata, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	query := tx

	var totalItem int64
	if err := query.Model(&entity.PaymentMethod{}).Count(&totalItem).Error; err != nil {
		c.Log.WithError(err).Error("Failed to count payment methods")
		return nil, nil, err
	}

	if request.Page > 0 && request.Limit > 0 {
		offset := (request.Page - 1) * request.Limit
		query = query.Offset(offset).Limit(request.Limit)
	}

	paymentMethods, err := c.PaymentMethodRespository.FindAll(query)
	if err != nil {
		c.Log.WithError(err).Error("Failed to find payment methods")
		return nil, nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, nil, err
	}

	response := make([]*models.PaymentMethodResponse, len(paymentMethods))
	for i, pm := range paymentMethods {
		response[i] = converter.PaymentMethodtoResponse(&pm)
	}

	totalPage := 1
	if request.Limit > 0 {
		totalPage = int((totalItem + int64(request.Limit) - 1) / int64(request.Limit))
	}

	paging := &models.PageMetadata{
		Page:      request.Page,
		Limit:     request.Limit,
		TotalItem: int(totalItem),
		TotalPage: totalPage,
	}

	return response, paging, nil
}
