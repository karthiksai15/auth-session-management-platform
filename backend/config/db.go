package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver — the underscore means we only need its side effects (registering the driver)
)

// DB is the global database connection pool.
// Other packages will use this to run queries.
var DB *sql.DB

// ConnectDB reads the DB credentials from environment variables,
// opens a connection to PostgreSQL, and checks that it actually works.
func ConnectDB() {
	// Read database connection details from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Build the PostgreSQL connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	// Open the connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open DB connection: ", err)
	}

	// Ping to verify the connection is actually working
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB: ", err)
	}

	DB = db
	log.Println("Connected to PostgreSQL database successfully")
}
