package database

import (
	"log"

	"github.com/Nohty/api/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Panicln("failed to connect to database")
	}

	log.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{}, &model.Address{}, &model.Delivery{})
	log.Println("Database Migrated")
}
