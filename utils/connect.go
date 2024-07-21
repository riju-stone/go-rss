package utils

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	log "github.com/riju-stone/go-rss/logging"
)

func ConnectDB() *sql.DB {
	dbString := os.Getenv("DATABASE_URI")
	if dbString == "" {
		log.Panic("DB connection string not found")
	}

	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Panic("DB connection failed with error: %s", err)
	}

	var dbVersion string
	if err := db.QueryRow("select version()").Scan(&dbVersion); err != nil {
		log.Panic("DB connection failed with error: %s", err)
	}

	log.Info("DB Connected")
	log.Info("DB Version: %s", dbVersion)

	return db
}
