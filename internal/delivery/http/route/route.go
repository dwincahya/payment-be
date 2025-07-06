package route

import (
	"github.com/dwincahya/payment-be/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                      *fiber.App
	PaymentChannelController *http.PaymentChannelController
	PaymentMethodController  *http.PaymentMethodController
}

func (c *RouteConfig) Setup() {
	c.App.Get("/api/channels", c.PaymentChannelController.List)
	c.App.Post("/api/channels", c.PaymentChannelController.Create)
	c.App.Put("/api/channels/:id", c.PaymentChannelController.Update)
	c.App.Get("/api/channels/:id", c.PaymentChannelController.Get)
	c.App.Delete("/api/channels/:id", c.PaymentChannelController.Delete)

	c.App.Get("/api/methods", c.PaymentMethodController.List)
	c.App.Post("/api/methods", c.PaymentMethodController.Create)
	c.App.Put("/api/methods/:id", c.PaymentMethodController.Update)
	c.App.Get("/api/methods/:id", c.PaymentMethodController.Get)
	c.App.Delete("/api/methods/:id", c.PaymentMethodController.Delete)
}
