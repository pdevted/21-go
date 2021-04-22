package database

import (
	"fmt"

	// TODO: gorm v2
	"github.com/jinzhu/gorm" // gorm v1
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")

	// dbUsername := os.Getenv("DB_USERNAME")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_TABLE")
	// dbPort := os.Getenv("DB_PORT")
	dbUsername := "taedongkim"
	dbPassword := "12345"
	dbHost := "localhost"
	dbName := "gorm"
	dbPort := "5433"

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbName, dbPassword)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err == nil {
		return db, err
	}

	return nil, nil
}
