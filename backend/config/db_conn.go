package config

/*
	Create a connection to the database
*/

import (
	"backend/internal/domain"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	log.Println("Connecting to database...")

	dbUrl := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(
		&domain.User{},
		&domain.Hotel{},
		&domain.Room{},
		&domain.Facility{},
		&domain.Booking{},
	)
	log.Println("Database connected successfully")
	return db
}
