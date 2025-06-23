package routes

import (
	"github.com/dwincahya/payment-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterChannelRoutes(app *fiber.App) {
	channel := app.Group("/channels")

	channel.Get("/", controllers.GetAllChannels)
	channel.Get("/:id", controllers.GetChannelByID)
	channel.Post("/", controllers.CreateChannel)
	channel.Put("/:id", controllers.UpdateChannel)
	channel.Delete("/:id", controllers.DeleteChannel)
}
