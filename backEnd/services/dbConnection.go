package services

import (
	"github.com/golangBaseProject/backEnd/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func OpenDatabase() {
	dsn := "host=localhost user=postgres password=007009 dbname=example port=5432 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Db = connection
	connection.AutoMigrate(&model.Users{})
}
