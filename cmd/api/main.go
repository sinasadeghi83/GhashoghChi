package main

import (
	"database/sql"
	"log"

	"github.com/sinasadeghi83/ghashoghchi/internal/api/server"
	"github.com/sinasadeghi83/ghashoghchi/internal/config"
	"github.com/sinasadeghi83/ghashoghchi/internal/platform/database"
)

func main() {
	cfg := config.LoadConfig()

	//Connect to DB
	db, err := database.OpenDatabase(cfg.DbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get generic database object: %v", sqlDb)
	}

	defer closeDB(sqlDb)

	srv := server.NewServer(cfg.ServerPort, db)
	srv.SetupRoutes()

	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

func closeDB(sqlDb *sql.DB) {
	if err := sqlDb.Close(); err != nil {
		log.Printf("Failed to close database connection: %v", err)
	}
}
