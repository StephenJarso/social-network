package main

import (
	"log"

	"backend/pkg/db"
)

func main() {
	// -----------------------------
	// Application configuration
	// -----------------------------

	dbConfig := db.DefaultConfig("data/app.db")

	//serverAddr := ":8080"

	// -----------------------------
	// Initialize database
	// -----------------------------

	database, err := db.NewSQLite(dbConfig)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer database.Close()

	log.Println("database connected")
}