package routes

import (
	"github.com/dwincahya/payment-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterMethodRoutes(app *fiber.App) {
	method := app.Group("/api/methods")
	method.Get("/", controllers.GetAllMethods)
	method.Get("/:id", controllers.GetMethodByID)
	method.Post("/", controllers.CreateMethod)
	method.Put("/:id", controllers.UpdateMethod)
	method.Delete("/:id", controllers.DeleteMethod)
}
