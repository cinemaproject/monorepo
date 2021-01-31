package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cinematic/monorepo/backend/pkg"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	dbURL := os.Getenv("BACKEND_DB_URL")

	var db *sqlx.DB
	var err error

	for true {
		db, err = sqlx.Connect("pgx", dbURL)
		if err == nil {
			break
		} else {
			log.Printf("Failed to connect to db: %s", err)
			time.Sleep(3*time.Second)
		}
	}

	router := pkg.InitializeRouter(db)

	err = http.ListenAndServe("0.0.0.0:3000", router)
	if err != nil {
		log.Fatalf("Failed to listen to server at 0.0.0.0:3000: %s", err)
	}
}
