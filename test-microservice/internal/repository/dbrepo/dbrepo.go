package dbrepo

import "test-microservice/internal/repository"

type postgresDBRepo struct {
}

type inMemoryDBRepo struct {
}

func NewPostgresDBRepo() repository.DatabaseRepo {
	return &postgresDBRepo{}
}

func NewInMemoryDBRepo() repository.DatabaseRepo {
	return inMemoryDBRepo{}
}
