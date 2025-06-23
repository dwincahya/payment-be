package controllers

import (
	"github.com/dwincahya/payment-be/database"
	"github.com/dwincahya/payment-be/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllChannels(c *fiber.Ctx) error {
	var channels []models.PaymentChannel
	database.DB.Order("id ASC").Find(&channels)
	return c.JSON(channels)
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
