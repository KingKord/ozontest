package dbrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"log"
	"time"
)

func (p postgresDBRepo) Add(URL string, shortURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into urls (shorturl, longurl) values ($1, $2)`
	_, err := p.DB.ExecContext(ctx, query,
		shortURL,
		URL,
	)
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok {
			log.Println("Duplicate hashes:", err)
			return pgErr
		}

		// Handle other errors
		log.Println("Other error:", err)
		return err
	}

	return nil
}

// GetShortURL retrieves short URL from DB via long URL. If there is no long URL in DB returns "" and nil
func (p postgresDBRepo) GetShortURL(URL string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select shorturl from urls where longurl = $1`

	var shortURL string
	err := p.DB.QueryRowContext(ctx, query, URL).Scan(&shortURL)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return "", nil
	case err != nil:
		return "", err
	default:
		break
	}

	return shortURL, nil
}

func (p postgresDBRepo) GetLongURL(shortURL string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select longurl from urls where shorturl = $1`

	var longURL string
	err := p.DB.QueryRowContext(ctx, query, shortURL).Scan(&longURL)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return "", errors.New(fmt.Sprintf("no such short URL, %v", err))
	case err != nil:
		return "", err
	default:
		break
	}

	return longURL, nil
}
