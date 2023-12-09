package driver

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"os"
	"time"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	var counts int

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			migrateUp()
			return connection
		}

		if counts > 10 {
			log.Fatal(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(time.Second * 2)
		continue
	}
}

func migrateUp() {
	dbURL := os.Getenv("dbURL")
	migrationPath := os.Getenv("migrationPath")
	if migrationPath == "" || dbURL == "" {
		log.Println("no env variables...")
		return
	}

	// Create a new instance of migrate.
	m, err := migrate.New(migrationPath, dbURL)
	if err != nil {
		log.Println("Error creating migrate instance:", err)
		os.Exit(1)
	}

	// Migrations to the last version.
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Println("Error applying migrations:", err)
		os.Exit(1)
	}

	log.Println("Migrations applied successfully.")
}
