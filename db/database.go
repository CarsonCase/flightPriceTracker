package db

// db/database.go

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// DB is the database connection instance.
var DB *sql.DB

// InitDB initializes the database connection.
func InitDB() error {
	godotenv.Load()
	DB_STRING := os.Getenv("DB_URL")
	// Configure the PostgreSQL connection string.
	connStr := DB_STRING

	// Open a connection to the database.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// Verify the connection to the database.
	if err = db.Ping(); err != nil {
		return err
	}

	// Set the DB variable to the opened connection.
	DB = db

	fmt.Println("Database connected successfully")
	return nil
}
