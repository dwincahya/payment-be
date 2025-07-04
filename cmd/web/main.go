package main

import (
	"fmt"

	"github.com/dwincahya/payment-be/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validate := config.NewValidator(viperConfig)
	app := config.NewFiber(viperConfig)

	appConfig := &config.AppConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Viper:    viperConfig,
	}

	config.Bootstrap(appConfig)

	webPort := viperConfig.GetInt("WEB_PORT")
	if err := app.Listen(fmt.Sprintf(":%d", webPort)); err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}
