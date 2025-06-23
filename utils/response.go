package utils

import "github.com/gofiber/fiber/v2"

func JSONError(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(fiber.Map{"error": msg})
}

func JSONSuccess(c *fiber.Ctx, data interface{}) error {
	return c.JSON(data)
}
