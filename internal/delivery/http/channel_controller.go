package http

import (
	"errors"
	"strconv"

	"github.com/dwincahya/payment-be/internal/helper"
	models "github.com/dwincahya/payment-be/internal/model"
	"github.com/dwincahya/payment-be/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PaymentChannelController struct {
	UseCase  *usecase.PaymentChannelUseCase
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewPaymentChannelController(useCase *usecase.PaymentChannelUseCase, log *logrus.Logger) *PaymentChannelController {
	return &PaymentChannelController{
		UseCase:  useCase,
		Log:      log,
		Validate: validator.New(),
	}
}

func (c *PaymentChannelController) Create(ctx *fiber.Ctx) error {
	request := new(models.CreatePaymentChannelRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse body")
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("Validation failed")
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	response, err := c.UseCase.Create(ctx.Context(), request)
	if err != nil {
		return c.handleError(ctx, err)
	}

	return helper.SuccessResponse(ctx, response, "Payment channel created successfully")

}

func (c *PaymentChannelController) List(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	var paymentMethodID *uint
	if param := ctx.Query("payment_method_id"); param != "" {
		idConv, err := strconv.ParseUint(param, 10, 32)
		if err == nil {
			tmp := uint(idConv)
			paymentMethodID = &tmp
		}
	}

	request := &models.ListPaymentChannelRequest{
		PaymentMethodID: paymentMethodID,
		Page:            page,
		Limit:           limit,
	}

	data, paging, err := c.UseCase.List(ctx.Context(), request)
	if err != nil {
		return c.handleError(ctx, err)
	}

	return helper.SuccessResponseWithPaging(ctx, data, "Payment channel list", paging)
}

func (c *PaymentChannelController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.Log.WithError(err).Error("invalid id param")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid id param")
	}

	request := &models.GetPaymentChannelRequest{
		ID: uint(uintID),
	}

	response, err := c.UseCase.Get(ctx.Context(), request)
	if err != nil {
		return c.handleError(ctx, err)
	}

	return helper.SuccessResponse(ctx, response, "Payment channel GET Success")
}

func (c *PaymentChannelController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	request := new(models.UpdatePaymentChannelRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse body")
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("Validation failed")
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.Log.WithError(err).Error("invalid id param")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid id param")
	}
	request.ID = uint(uintID)

	response, err := c.UseCase.Update(ctx.Context(), request)
	if err != nil {
		return c.handleError(ctx, err)
	}

	return helper.SuccessResponse(ctx, response, "Payment channel update success")
}

func (c *PaymentChannelController) Delete(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	if idParam == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	idUint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid id")
	}
	id := uint(idUint64)

	request := &models.DeletePaymentChannelRequest{
		ID: id,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		return c.handleError(ctx, err)
	}

	return helper.SuccessResponse(ctx, true, "Payment channel deleted")
}

func (c *PaymentChannelController) handleError(ctx *fiber.Ctx, err error) error {
	c.Log.WithError(err).Error("Request failed")

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		var validationDetails []string
		for _, fieldErr := range ve {
			validationDetails = append(validationDetails, fieldErr.Field()+": "+fieldErr.ActualTag())
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  validationDetails,
		})
	}

	if err.Error() == "payment channel with this code already exists" {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": err.Error(),
	})
}
