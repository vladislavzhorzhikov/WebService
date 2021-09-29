package db

import (
	"WebService/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func MigrateDB() {
	dsn := "host=localhost user=zvm password=zvm dbname=webservice port=5432 sslmode=disable"
	var err error
	DataBase, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DataBase.AutoMigrate(&models.Worker{}, &models.Passport{}, &models.Department{})
}
