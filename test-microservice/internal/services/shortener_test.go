package services

import (
	"github.com/stretchr/testify/assert"
	"test-microservice/internal/repository/dbrepo"
	"testing"
)

var testsForShortenURL = []struct {
	num           int
	description   string
	longURL       string
	isErrExpected bool
	errText       string
	expected      string
}{
	{1, "all fine 1", "youtube.com/channels/beast?redirect=true&from=finder", false, "", ""},
	{2, "case where we already have short URL in DB", "test.com/existing", false, "", "existingHash"},
	{3, "case where hash got duplicated twice", "test.com/duplicate-hash", true, "duplicate hashes", ""},
	{4, "case where error occurred during paste in DB", "test.com/err-in-DB", true, "error in DB", ""},
	{5, "all fine", "youtube.com/channels/beast?redirect=true&from=browser", false, "", ""},
}

func TestURLShortenerByRandomizing_ShortenURL(t *testing.T) {
	mockRepo := dbrepo.NewTestingDBRepo() // Используем мок для базы данных
	urlShortener := NewURLShortenerByRandomizing(mockRepo)

	for _, test := range testsForShortenURL {
		shortURL, err := urlShortener.ShortenURL(test.longURL)
		if test.isErrExpected {
			assert.Error(t, err)
			assert.Contains(t, err.Error(), test.errText)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, shortURL)
			if test.expected != "" {
				assert.Contains(t, shortURL, test.expected)
			}
		}
	}
}

var testsForGetOriginalURL = []struct {
	num           int
	description   string
	shortURL      string
	isErrExpected bool
	errText       string
	expected      string
}{
	{1, "all fine", "existingHash", false, "", "youtube.com/channels/beast?redirect=true&from=finder"},
	{2, "case where hash does not exist", "non-existenceHash", true, "hash not exist in DB", ""},
	{3, "case where error occurred in DB", "invalidDB", true, "invalid DB", ""},
}

func TestURLShortenerByRandomizing_GetOriginalURL(t *testing.T) {
	mockRepo := dbrepo.NewTestingDBRepo() // Используем мок для базы данных
	urlShortener := NewURLShortenerByRandomizing(mockRepo)

	for _, test := range testsForGetOriginalURL {
		originalURL, err := urlShortener.GetOriginalURL(test.shortURL)
		if test.isErrExpected {
			assert.Error(t, err)
			assert.Contains(t, err.Error(), test.errText)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, originalURL)
			if test.expected != "" {
				assert.Contains(t, originalURL, test.expected)
			}
		}
	}
}
