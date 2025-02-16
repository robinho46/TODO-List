package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "todo-db" // Replace with your actual PostgreSQL password
	dbname   = "todo_list"
)

var db *sql.DB

func connectDB() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("❌ Error connecting to the database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("❌ Database is not responding:", err)
	}

	fmt.Println("✅ Successfully connected to the database!")
}
