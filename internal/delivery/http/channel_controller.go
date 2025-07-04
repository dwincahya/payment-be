package http

import (
	"strconv"

	models "github.com/dwincahya/payment-be/internal/model"
	"github.com/dwincahya/payment-be/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PaymentChannelController struct {
	UseCase *usecase.PaymentChannelUseCase
	Log     *logrus.Logger
}

func NewPaymentChannelController(useCase *usecase.PaymentChannelUseCase, log *logrus.Logger) *PaymentChannelController {
	return &PaymentChannelController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *PaymentChannelController) Create(ctx *fiber.Ctx) error {
	request := new(models.CreatePaymentChannelRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to create payment channel")
		return err
	}

	return ctx.JSON(models.WebResponse[*models.PaymentChannelResponse]{Data: response})
}

func (c *PaymentChannelController) List(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		page = 1
	}

	size, err := strconv.Atoi(ctx.Query("size", "10"))
	if err != nil {
		size = 10
	}

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
		Limit:           size,
	}

	response, err := c.UseCase.List(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to list payment channels")
		return err
	}

	return ctx.JSON(models.WebResponse[[]*models.PaymentChannelResponse]{Data: response})
}

func (c *PaymentChannelController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.Log.WithError(err).Error("invalid id param")
		return fiber.ErrBadRequest
	}

	request := &models.GetPaymentChannelRequest{
		ID: uint(uintID),
	}

	response, err := c.UseCase.Get(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to get payment channel")
		return err
	}

	return ctx.JSON(models.WebResponse[*models.PaymentChannelResponse]{Data: response})
}

func (c *PaymentChannelController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	request := new(models.UpdatePaymentChannelRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse body")
		return fiber.ErrBadRequest
	}

	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.Log.WithError(err).Error("invalid id param")
		return fiber.ErrBadRequest
	}
	request.ID = uint(uintID)

	response, err := c.UseCase.Update(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to update payment channel")
		return err
	}

	return ctx.JSON(models.WebResponse[*models.PaymentChannelResponse]{Data: response})
}

func (c *PaymentChannelController) Delete(ctx *fiber.Ctx) error {

	idParam := ctx.Params("id")
	if idParam == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	idUint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	id := uint(idUint64)

	request := &models.DeletePaymentChannelRequest{
		ID: id,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete paymentchannel")
		return err
	}

	return ctx.JSON(models.WebResponse[bool]{Data: true})
}
