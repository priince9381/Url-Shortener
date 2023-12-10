package database

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Database struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Dbname   string `toml:"dbname"`
	SSLMode  string `toml:"sslmode"`
}

func InitDB(ctx context.Context, config Database) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.Dbname, config.SSLMode)

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to the database")

	// Auto-create tables and schema if not exists
	err = DB.AutoMigrate(&URL{})
	if err != nil {
		log.Fatal("Failed to auto migrate tables:", err)
	}
}
