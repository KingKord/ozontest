package services

import (
	"crypto/rand"
	"math/big"
	"test-microservice/internal/repository"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	length  = 10
)

type URLShortener interface {
	ShortenURL(URL string) (string, error)
	GetOriginalURL(shortURL string) (string, error)
}

type URLShortenerByRandomizing struct {
	repo repository.DatabaseRepo
}

func NewURLShortener(databaseRepo repository.DatabaseRepo) URLShortener {
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
	shortURL, err := generateShortLink()
	if err != nil {
		return "", err
	}

	err = U.repo.Add(URL, shortURL)
	if err != nil {
		return "", err
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
