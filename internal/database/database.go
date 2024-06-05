package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/yhsgw178/todoapp/internal/config"
)

var DB *sql.DB

func InitDB(cfg *config.DBConfig) {
	var err error
	DB, err = sql.Open("postgres", cfg.DSN())
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Could not ping the database: %v", err)
	}

	log.Println("Connected to the database", cfg.DSN())
}
