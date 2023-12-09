package inits

import (
	"os"

	models "github.com/Joybaruah/ssl-tracker/pkg/model"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	DBMigration()
}

func DBMigration() {
	DB.AutoMigrate(&models.User{})
}
