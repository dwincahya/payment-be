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

func SuccessResponse(ctx *fiber.Ctx, data interface{}, message string) error {
	response := models.WebResponse[interface{}]{
		Code:    200,
		Data:    data,
		Message: message,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func SuccessResponseWithPaging(ctx *fiber.Ctx, data interface{}, message string, paging *models.PageMetadata) error {
	response := models.WebResponse[interface{}]{
		Code:    200,
		Data:    data,
		Paging:  paging,
		Message: message,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
