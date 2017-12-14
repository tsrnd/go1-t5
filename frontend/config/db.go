package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	db "github.com/tsrnd/goweb5/frontend/services/database/sql"
)

// DB func
func DB() *sql.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbDlct := os.Getenv("DATABASE_DLCT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	db, err := db.Connect(dbDlct, dbUser, dbPass, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
