package helper

import (
	models "github.com/dwincahya/payment-be/internal/model"
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(ctx *fiber.Ctx, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(models.WebResponse[any]{
		Code:    statusCode,
		Message: message,
		Data:    nil,
	})
}

func SuccessResponse[T any](ctx *fiber.Ctx, data T) error {
	return ctx.JSON(models.WebResponse[T]{
		Code:    200,
		Message: "Success",
		Data:    data,
	})
}
