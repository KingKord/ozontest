package main

import (
	"os"
	"test-microservice/internal/config"
	"test-microservice/internal/driver"
	"test-microservice/internal/repository"
	"test-microservice/internal/services"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	gRPCPort = "50001"
)

var app config.Config

func main() {
	dbType := os.Getenv("dbType")
	var dbRepo repository.DatabaseRepo

	if dbType == "postgres" {
		db := driver.ConnectToDB()
		dbRepo = repository.NewPostgresDBRepo(db)
	} else {
		dbRepo = repository.NewInMemoryDBRepo()
	}

	testService := services.NewURLShortenerByRandomizing(dbRepo)
	app = config.Config{
		DB:      dbRepo,
		Service: testService,
	}

	gRPCListen()
}
