package database

import (
	"log"

	"github.com/spectertian/dianputest/config"
	"github.com/spectertian/dianputest/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err = DB.AutoMigrate(&model.User{}, &model.Shop{}); err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	log.Println("Database connected and migrated successfully")
}
