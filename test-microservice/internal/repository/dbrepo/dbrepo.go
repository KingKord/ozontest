package dbrepo

import (
	"database/sql"
	"sync"
	"test-microservice/internal/repository"
)

type postgresDBRepo struct {
	DB *sql.DB
}

type inMemoryDBRepo struct {
	shortURLs map[string]string
	longURLs  map[string]string
	mu        sync.RWMutex
}

func NewPostgresDBRepo(conn *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		DB: conn,
	}
}

func NewInMemoryDBRepo() repository.DatabaseRepo {
	return &inMemoryDBRepo{
		shortURLs: make(map[string]string),
		longURLs:  make(map[string]string),
		mu:        sync.RWMutex{},
	}
}
