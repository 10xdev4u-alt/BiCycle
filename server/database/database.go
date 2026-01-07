package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/princetheprogrammerbtw/TheBiCycleApp/server/config"
)

// InitDB initializes and returns a database connection
func InitDB(cfg *config.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Successfully connected to the database!")
	return db
}
