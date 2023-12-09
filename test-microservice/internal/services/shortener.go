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

func (u URLShortenerByRandomizing) ShortenURL(URL string) (string, error) {
	shortURLFromDB, err := u.repo.GetShortURL(URL)
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

	err = u.repo.Add(URL, shortURL)
	if err != nil {
		// Handle err when hashes duplicates with already existing
		if err.Error() == "duplicate hashes" {
			shortURL, err = generateShortLink()
			if err != nil {
				return "", err
			}
			err = u.repo.Add(URL, shortURL)
		}
		return "", err
	}

	return shortURL, nil
}

func (u URLShortenerByRandomizing) GetOriginalURL(shortURL string) (string, error) {
	longURL, err := u.repo.GetLongURL(shortURL)
	if err != nil {
		return "", err
	}

	return longURL, nil
}
