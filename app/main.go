package main

import (
	"github.com/dwincahya/payment-be/config"
	"github.com/dwincahya/payment-be/database"
	"github.com/dwincahya/payment-be/models"
	"github.com/dwincahya/payment-be/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.LoadEnv()
	database.Connect()

	database.DB.AutoMigrate(&models.PaymentMethod{}, &models.PaymentChannel{})

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))
	routes.RegisterMethodRoutes(app)
	routes.RegisterChannelRoutes(app)

	app.Listen(":" + config.GetEnv("PORT", "3000"))
}
