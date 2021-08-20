package services

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/golangBaseProject/model"
  )
  
var Db *gorm.DB

func OpenDatabase(){
	dsn := "host=localhost user=postgres password=postgres dbname=example port=5432 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Db = connection
	connection.AutoMigrate(&model.Users{})
}
