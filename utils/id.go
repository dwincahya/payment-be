package utils

import (
	"github.com/dwincahya/payment-be/database"
	"github.com/dwincahya/payment-be/models"
)

func GetNextMethodID() uint {
	var last models.PaymentMethod
	database.DB.Order("id desc").First(&last)
	return last.ID + 1
}

func GetNextChannelID() uint {
	var last models.PaymentChannel
	database.DB.Order("id desc").First(&last)
	return last.ID + 1
}
