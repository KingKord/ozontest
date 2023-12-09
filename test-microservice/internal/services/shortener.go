package services

import (
	"crypto/rand"
	"github.com/jackc/pgconn"
	"math/big"
	"test-microservice/internal/repository"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	length  = 10
)

//var app config.Config

type URLShortener interface {
	ShortenURL(URL string) (string, error)
	GetOriginalURL(shortURL string) (string, error)
}

type URLShortenerByRandomizing struct {
	repo repository.DatabaseRepo
}

func NewURLShortenerByRandomizing(databaseRepo repository.DatabaseRepo) URLShortener {
	return &URLShortenerByRandomizing{
		repo: databaseRepo,
	}
}

func generateShortLink() (string, error) {
	var shortLink []byte
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		shortLink = append(shortLink, letters[num.Int64()])
	}
	return string(shortLink), nil
}

func (U URLShortenerByRandomizing) ShortenURL(URL string) (string, error) {
	shortURLFromDB, err := U.repo.GetShortURL(URL)
	if err != nil {
		return "", err
	}
	if shortURLFromDB != "" {
		return shortURLFromDB, nil
	}

	shortURL, err := generateShortLink()
	if err != nil {
		return "", err

	}

	err = U.repo.Add(URL, shortURL)
	if err != nil {
		//if IsPostgres {
		_, ok := err.(*pgconn.PgError)
		if ok { // Handle err when hashes duplicates with already existing
			shortURL, err = generateShortLink()
			if err != nil {
				return "", err
			}
			err = U.repo.Add(URL, shortURL)
		}

		return "", err
		//}
		//return "", err
	}

	return shortURL, nil
}

func (U URLShortenerByRandomizing) GetOriginalURL(shortURL string) (string, error) {
	longURL, err := U.repo.GetLongURL(shortURL)
	if err != nil {
		return "", err
	}

	return longURL, nil
}
