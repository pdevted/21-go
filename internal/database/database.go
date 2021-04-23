package database

import (
	"fmt"

	// TODO: gorm v2
	"github.com/jinzhu/gorm" // gorm v1
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

func NewDatabase() (*gorm.DB, error) {
	log.Info("Setting up new database connection")

	// dbUsername := os.Getenv("DB_USERNAME")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_TABLE")
	// dbPort := os.Getenv("DB_PORT")
	// sslMode := os.Getenv("SSL_MODE")
	dbUsername := "taedongkim"
	dbPassword := "12345"
	dbHost := "localhost"
	dbName := "gorm"
	dbPort := "5433"
	sslMode := "disable" // or require

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUsername, dbName, dbPassword, sslMode)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err == nil {
		return db, err
	}

	return nil, nil
}
