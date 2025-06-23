package controllers

import (
	"github.com/dwincahya/payment-be/database"
	"github.com/dwincahya/payment-be/models"
	"github.com/dwincahya/payment-be/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAllChannels(c *fiber.Ctx) error {
	page := c.Query("page")
	limit := c.Query("limit")

	var channels []models.PaymentChannel

	if page == "" && limit == "" {
		if err := database.DB.Order("id ASC").Find(&channels).Error; err != nil {
			return utils.JSONError(c, 500, "Failed to fetch channels")
		}
		return c.JSON(channels)
	}

	paginate := utils.ParsePaginationParams(c)

	var total int64
	query := database.DB.Model(&models.PaymentChannel{})
	query.Count(&total)

	err := query.
		Order("id ASC").
		Limit(paginate.Limit).
		Offset(paginate.Skip).
		Find(&channels).Error

	if err != nil {
		return utils.JSONError(c, 500, "Failed to fetch channels")
	}

	return c.JSON(fiber.Map{
		"data":       channels,
		"page":       paginate.Page,
		"limit":      paginate.Limit,
		"total":      total,
		"totalPages": (total + int64(paginate.Limit) - 1) / int64(paginate.Limit),
	})
}

func GetChannelByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var channel models.PaymentChannel
	if err := database.DB.First(&channel, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Channel not found"})
	}
	return c.JSON(channel)
}

func CreateChannel(c *fiber.Ctx) error {
	var channel models.PaymentChannel
	if err := c.BodyParser(&channel); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := database.DB.Create(&channel).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create channel"})
	}
	return c.Status(201).JSON(channel)
}

func UpdateChannel(c *fiber.Ctx) error {
	id := c.Params("id")
	var channel models.PaymentChannel
	if err := database.DB.First(&channel, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Channel not found"})
	}
	if err := c.BodyParser(&channel); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	database.DB.Save(&channel)
	return c.JSON(channel)
}

func DeleteChannel(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.PaymentChannel{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete channel"})
	}
	return c.SendStatus(204)
}
