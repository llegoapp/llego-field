package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var dbPool *sql.DB

// InitPool initializes the database connection pool.
func InitPool() error {
	var err error

	// Construct the database connection string
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Initialize the connection pool
	dbPool, err = sql.Open("pgx", connectionString)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Verify the connection
	if err = dbPool.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	return nil
}

// GetPool returns the database connection pool.
func GetPool() *sql.DB {
	return dbPool
}

// ClosePool closes the database connection pool.
func ClosePool() error {
	if dbPool != nil {
		return dbPool.Close()
	}
	return nil
}
