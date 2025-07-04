package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(v *viper.Viper, log *logrus.Logger) *gorm.DB {
	host := v.GetString("DB_HOST")
	port := v.GetString("DB_PORT")
	user := v.GetString("DB_USER")
	password := v.GetString("DB_PASSWORD")
	dbname := v.GetString("DB_NAME")
	idle := v.GetInt("DB_MAX_IDLE")
	max := v.GetInt("DB_MAX_CONN")
	lifetime := v.GetInt("DB_CONN_LIFETIME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get db instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(idle)
	sqlDB.SetMaxOpenConns(max)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(lifetime))

	return db
}
