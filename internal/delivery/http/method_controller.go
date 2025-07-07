package http

import (
	"strconv"

	"github.com/dwincahya/payment-be/internal/helper"
	models "github.com/dwincahya/payment-be/internal/model"
	"github.com/dwincahya/payment-be/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PaymentMethodController struct {
	UseCase  *usecase.PaymentMethodUseCase
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewPaymentMethodController(useCase *usecase.PaymentMethodUseCase, log *logrus.Logger) *PaymentMethodController {
	return &PaymentMethodController{
		UseCase:  useCase,
		Log:      log,
		Validate: validator.New(),
	}
}

func (c *PaymentMethodController) Create(ctx *fiber.Ctx) error {
	request := new(models.CreatePaymentMethodRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse body")
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("Validation failed")
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to create payment method")
		return err
	}

	return helper.SuccessResponse(ctx, response, "Payment method created succesfully")

}

func (c *PaymentMethodController) Update(ctx *fiber.Ctx) error {
	request := new(models.UpdatePaymentMethodRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse body")
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("Validation failed")
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	id, err := ctx.ParamsInt("id")
	if err != nil {
		c.Log.WithError(err).Error("Invalid payment method id")
		return fiber.ErrBadRequest
	}
	request.ID = uint(id)

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to update payment method")
		return err
	}

	return helper.SuccessResponse(ctx, response, "Payment method updated succesfully")
}

func (c *PaymentMethodController) Get(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		c.Log.WithError(err).Error("Invalid payment method id")
		return fiber.ErrBadRequest
	}

	request := &models.GetPaymentMethodRequest{
		ID: uint(id),
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to get payment method")
		return err
	}

	return helper.SuccessResponse(ctx, response, "Payment method get")
}

func (c *PaymentMethodController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		c.Log.WithError(err).Error("Invalid payment method id")
		return fiber.ErrBadRequest
	}

	request := &models.DeletePaymentMethodRequest{
		ID: uint(id),
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("Failed to delete payment method")
		return err
	}

	return helper.SuccessResponse(ctx, true, "Payment method deleted")
}

func (c *PaymentMethodController) List(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	request := &models.ListPaymentMethodRequest{
		Page:  page,
		Limit: limit,
	}

	data, paging, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to list payment methods")
		return err
	}

	return helper.SuccessResponseWithPaging(ctx, data, "Payment method list", paging)
}
