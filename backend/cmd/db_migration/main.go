package main

import (
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

func insertStubData(db *sqlx.DB) {
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO films VALUES (1, 'Sample Film', 'sample_url', 'movie', 2019, 0, 100, 't000000001', 'Lorem ipsum dolor sit amet', '6284065');")
	tx.MustExec("INSERT INTO people VALUES (1, 'John Doe', 'http://example.com', 1985, 0, 'nm000000000001');")
	tx.MustExec("INSERT INTO relations VALUES (1, 1, 0);")
	tx.Commit()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	migrations := &migrate.FileMigrationSource{
		Dir: "schema",
	}

	dbURL := os.Getenv("BACKEND_DB_URL")
	env := os.Getenv("BACKEND_ENV")

	var db *sqlx.DB
	var err error

	for true {
		db, err = sqlx.Connect("pgx", dbURL)
		if err == nil {
			break
		} else {
			log.Printf("Failed to connect to db: %s", err)
			time.Sleep(3 * time.Second)
		}
	}

	defer db.Close()

	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	log.Printf("Applied %d migrations!\n", n)
	if err != nil {
		panic(err)
	}

	if env == "dev" {
		insertStubData(db)
	}
}
