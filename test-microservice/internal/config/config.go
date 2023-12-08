package config

import (
	"test-microservice/internal/repository"
	"test-microservice/internal/services"
)

type Config struct {
	DB      repository.DatabaseRepo
	Service services.URLShortener
}
