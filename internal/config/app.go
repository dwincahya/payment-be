package config

import (
	"github.com/dwincahya/payment-be/internal/delivery/http"
	"github.com/dwincahya/payment-be/internal/delivery/http/route"
	"github.com/dwincahya/payment-be/internal/repository"
	"github.com/dwincahya/payment-be/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppConfig struct {
	App      *fiber.App
	DB       *gorm.DB
	Log      *logrus.Logger
	Validate *validator.Validate
	Viper    *viper.Viper
}

func Bootstrap(config *AppConfig) {
	paymentMethodRepo := repository.NewPaymentMethodRepository(config.Log)
	paymentChannelRepo := repository.NewPaymentChannelRepository(config.Log)

	paymentMethodUC := usecase.NewPaymentMethodUseCase(config.DB, config.Log, config.Validate, paymentMethodRepo)
	paymentChannelUC := usecase.NewPaymentChannelUseCase(config.DB, config.Log, config.Validate, paymentChannelRepo, paymentMethodRepo)

	paymentMethodController := http.NewPaymentMethodController(paymentMethodUC, config.Log)
	paymentChannelController := http.NewPaymentChannelController(paymentChannelUC, config.Log)

	routeConfig := route.RouteConfig{
		App:                      config.App,
		PaymentChannelController: paymentChannelController,
		PaymentMethodController:  paymentMethodController,
	}

	routeConfig.Setup()
}
