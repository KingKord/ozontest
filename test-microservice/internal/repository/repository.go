package repository

import (
	"database/sql"
	"sync"
)

type DatabaseRepo interface {
	Add(URL string, shortURL string) error
	GetShortURL(URL string) (string, error)
	GetLongURL(shortURL string) (string, error)
}

type postgresDBRepo struct {
	db *sql.DB
}

type inMemoryDBRepo struct {
	shortURLs map[string]string
	longURLs  map[string]string
	mu        sync.RWMutex
}
