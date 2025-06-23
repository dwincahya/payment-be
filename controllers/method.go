package controllers

import (
	"github.com/dwincahya/payment-be/database"
	"github.com/dwincahya/payment-be/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllMethods(c *fiber.Ctx) error {
	var methods []models.PaymentMethod
	database.DB.Order("id ASC").Find(&methods)
	return c.JSON(methods)
}

func GetMethodByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var method models.PaymentMethod
	if err := database.DB.First(&method, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Method not found"})
	}
	return c.JSON(method)
}

func CreateMethod(c *fiber.Ctx) error {
	var method models.PaymentMethod
	if err := c.BodyParser(&method); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := database.DB.Create(&method).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create method"})
	}
	return c.Status(201).JSON(method)
}

func UpdateMethod(c *fiber.Ctx) error {
	id := c.Params("id")
	var method models.PaymentMethod
	if err := database.DB.First(&method, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Method not found"})
	}
	if err := c.BodyParser(&method); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	database.DB.Save(&method)
	return c.JSON(method)
}

func DeleteMethod(c *fiber.Ctx) error {
	id := c.Params("id")

	database.DB.Where("payment_method_id = ?", id).Delete(&models.PaymentChannel{})

	if err := database.DB.Delete(&models.PaymentMethod{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete method"})
	}
	return c.SendStatus(204)
}
