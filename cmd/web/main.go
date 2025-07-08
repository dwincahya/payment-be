package main

import (
	"fmt"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	_ "github.com/dwincahya/payment-be/docs"
	"github.com/dwincahya/payment-be/internal/config"
	"github.com/dwincahya/payment-be/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

// @title			SVD Payment API
// @version		1.0
// @description	This is a sample swagger for Fiber
// @host			103.210.35.189:8120
// @BasePath		/
func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validate := config.NewValidator(viperConfig)
	app := config.NewFiber(viperConfig)

	app.Use(middleware.NewCors())

	appConfig := &config.AppConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Viper:    viperConfig,
	}

	config.Bootstrap(appConfig)

	app.Get("/reference", func(ctx *fiber.Ctx) error {
		html, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL:  "./docs/swagger.json",
			DarkMode: true,
			Theme:    scalar.ThemeKepler,
			Layout:   scalar.LayoutModern,
		})
		if err != nil {
			log.Error("Failed to generate API reference HTML: " + err.Error())
			return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to generate API reference HTML")
		}
		return ctx.Type("html").SendString(html)
	})

	webPort := viperConfig.GetInt("WEB_PORT")
	if err := app.Listen(fmt.Sprintf(":%d", webPort)); err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}
