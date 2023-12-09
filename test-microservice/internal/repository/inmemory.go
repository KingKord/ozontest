package repository

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

func NewInMemoryDBRepo() DatabaseRepo {
	return &inMemoryDBRepo{
		shortURLs: make(map[string]string),
		longURLs:  make(map[string]string),
		mu:        sync.RWMutex{},
	}
}

func (i *inMemoryDBRepo) Add(URL string, shortURL string) error {
	log.Printf("adding %s, %s\n", URL, shortURL)

	i.mu.Lock()
	defer i.mu.Unlock()

	_, ok := i.shortURLs[shortURL]
	if ok {
		return errors.New("duplicate hashes")
	}

	_, ok = i.longURLs[URL]
	if ok {
		return errors.New("duplicate URLs")
	}

	i.shortURLs[shortURL] = URL
	i.longURLs[URL] = shortURL

	return nil
}

func (i *inMemoryDBRepo) GetShortURL(longURL string) (string, error) {
	log.Printf("get short URL, longURL: %s\n", longURL)
	i.mu.RLock()
	defer i.mu.RUnlock()

	shortURL := i.longURLs[longURL]
	return shortURL, nil
}

func (i *inMemoryDBRepo) GetLongURL(shortURL string) (string, error) {
	log.Printf("get long URL, shortURL: %s\n", shortURL)

	i.mu.RLock()
	defer i.mu.RUnlock()

	longURL, exists := i.shortURLs[shortURL]
	if !exists {
		return "", fmt.Errorf("shortURL not found")
	}

	return longURL, nil
}
