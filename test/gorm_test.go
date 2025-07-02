package test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormConnection() *gorm.DB {
	dialect := "host=localhost user=postgres password=@Ainosi2025! dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dialect), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect database", err)
	}

	return db

}

var db = GormConnection()

func TestConnection(t *testing.T) {
	assert.NotNil(t, db)
}
