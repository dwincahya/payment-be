package controllers

import (
	"github.com/dwincahya/payment-be/database"
	"github.com/dwincahya/payment-be/models"
	"github.com/dwincahya/payment-be/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAllMethods(c *fiber.Ctx) error {
	page := c.Query("page")
	limit := c.Query("limit")

	var methods []models.PaymentMethod

	if page == "" && limit == "" {
		if err := database.DB.Order("id ASC").Find(&methods).Error; err != nil {
			return utils.JSONError(c, 500, "Failed to fetch methods")
		}
		return c.JSON(methods)
	}

	paginate := utils.ParsePaginationParams(c)

	var total int64
	query := database.DB.Model(&models.PaymentMethod{})
	query.Count(&total)

	err := query.
		Order("id ASC").
		Limit(paginate.Limit).
		Offset(paginate.Skip).
		Find(&methods).Error

	if err != nil {
		return utils.JSONError(c, 500, "Failed to fetch methods")
	}

	return c.JSON(fiber.Map{
		"data":       methods,
		"page":       paginate.Page,
		"limit":      paginate.Limit,
		"total":      total,
		"totalPages": (total + int64(paginate.Limit) - 1) / int64(paginate.Limit),
	})
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
