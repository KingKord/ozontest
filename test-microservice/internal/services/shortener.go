package services

import "test-microservice/internal/repository"

type URLShortenerInterface interface {
}

type URLShortener struct {
	repo repository.DatabaseRepo
}

func NewURLShortener(databaseRepo repository.DatabaseRepo) URLShortenerInterface {
	return &URLShortener{
		repo: databaseRepo,
	}
}
