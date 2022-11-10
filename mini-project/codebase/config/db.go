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

type (
	dsn struct {
		Host     string
		Port     string
		User     string
		Password string
		Dbname   string
		SSLMode  string
		Timezone string
	}
)

func Database() {
	dsn := dsn{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
		Timezone: os.Getenv("DB_TIMEZONE"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	db_url := "host=" + dsn.Host + " user=" + dsn.User + " password=" + dsn.Password + " dbname=" + dsn.Dbname + " port=" + dsn.Port + " TimeZone=" + dsn.Timezone

	DB, err = gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")
}

func AutoMigrate() {
	DB.AutoMigrate(
		&entity.Category{},
		&entity.Product{},
		&entity.Cart{},
		&entity.Payment{},
	)
}
