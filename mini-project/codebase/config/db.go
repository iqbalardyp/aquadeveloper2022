package config

import (
	"ecommerce/codebase/entity"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Database() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
}

func AutoMigrate() {
	DB.AutoMigrate(
		&entity.Category{},
		&entity.Product{},
		&entity.Cart{},
		&entity.Payment{},
	)
}
