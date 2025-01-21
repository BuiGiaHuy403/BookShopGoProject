package testhelpers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	var (
		host     = "localhost"
		user     = "postgres"
		password = "12345"
		dbname   = "bookshoptest"
		port     = "5432"
	)
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func ConnectDBFailed() *gorm.DB {
	var (
		host     = "localhost"
		user     = "postgres"
		password = "1235" // Incorrect password to simulate failure
		dbname   = "bookshoptest"
		port     = "5432"
	)
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		fmt.Printf("Simulated connection failure: %v\n", err)
		return nil // Return nil for *gorm.DB
	}
	return db
}
