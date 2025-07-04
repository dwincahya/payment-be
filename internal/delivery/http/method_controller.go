package http

import (
	models "github.com/dwincahya/payment-be/internal/model"
	"github.com/dwincahya/payment-be/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PaymentMethodController struct {
	UseCase *usecase.PaymentMethodUseCase
	Log     *logrus.Logger
}

func NewPaymentMethodController(useCase *usecase.PaymentMethodUseCase, log *logrus.Logger) *PaymentMethodController {
	return &PaymentMethodController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *PaymentMethodController) Create(ctx *fiber.Ctx) error {
	request := new(models.CreatePaymentMethodRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse request body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to create payment method")
		return err
	}

	return ctx.JSON(models.WebResponse[*models.PaymentMethodResponse]{Data: response})
}

func (c *PaymentMethodController) Update(ctx *fiber.Ctx) error {
	request := new(models.UpdatePaymentMethodRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse request body")
		return fiber.ErrBadRequest
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

	return ctx.JSON(models.WebResponse[*models.PaymentMethodResponse]{Data: response})
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

	return ctx.JSON(models.WebResponse[*models.PaymentMethodResponse]{Data: response})
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

	return ctx.JSON(models.WebResponse[bool]{Data: true})
}

func (c *PaymentMethodController) List(ctx *fiber.Ctx) error {
	response, err := c.UseCase.List(ctx.UserContext())
	if err != nil {
		c.Log.WithError(err).Error("Failed to list payment methods")
		return err
	}

	return ctx.JSON(models.WebResponse[[]*models.PaymentMethodResponse]{Data: response})
}
