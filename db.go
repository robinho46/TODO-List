package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Global DB variable
var db *sql.DB

// connectDB initializes the connection to the Neon.tech database
func connectDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Retrieve the database URL from environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("❌ DATABASE_URL is not set in environment variables")
	}

	// Open the database connection
	db, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("❌ Error connecting to the database:", err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal("❌ Database is not responding:", err)
	}

	fmt.Println("✅ Successfully connected to the Neon.tech database!")
}

