package main

import (
	"github.com/cinematic/monorepo/backend/pkg"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
)

func main() {
	dbURL := os.Getenv("BACKEND_DB_URL")

	db, err := sqlx.Connect("pgx", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to db: %s", err)
	}

	router := pkg.InitializeRouter(db)

	err = http.ListenAndServe("0.0.0.0:3000", router)
	if err != nil {
		log.Fatalf("Failed to listen to server at 0.0.0.0:3000: %s", err)
	}
}
