package dbrepo

import (
	"database/sql"
	"test-microservice/internal/repository"
)

type postgresDBRepo struct {
	DB *sql.DB
}

type inMemoryDBRepo struct {
}

func NewPostgresDBRepo(conn *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		DB: conn,
	}
}

func NewInMemoryDBRepo() repository.DatabaseRepo {
	return inMemoryDBRepo{}
}
