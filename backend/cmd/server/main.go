package main

import (
	

	"backend/pkg/db"
)

func main() {
	// -----------------------------
	// Application configuration
	// -----------------------------

	dbConfig := db.Config{
		FilePath:        "data/app.db",
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
	}

	serverAddr := ":8080"

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