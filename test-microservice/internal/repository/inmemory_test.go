package repository

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewInMemoryDBRepo(t *testing.T) {
	repo := NewInMemoryDBRepo()

	if reflect.TypeOf(repo).String() != "*repository.inMemoryDBRepo" {
		t.Errorf("Did not get correct type from NewURLShortenerByRandomizing: got %s, wanted *dbrepo.inMemoryDBRepo", reflect.TypeOf(repo).String())
	}
}

var testsForAdd = []struct {
	num           int
	description   string
	URL           string
	shortURL      string
	isErrExpected bool
	errText       string
}{
	{1, "all fine", "test.com/all-fine", "hashed", false, ""},
	{2, "URL re-entry", "test.com/all-fine", "anotherHash", true, "duplicate URLs"},
	{3, "hash re-entry", "test.com/new-url", "hashed", true, "duplicate hashes"},
}

func TestInMemoryDBRepo_Add(t *testing.T) {
	repo := NewInMemoryDBRepo()

	for _, test := range testsForAdd {
		err := repo.Add(test.URL, test.shortURL)
		if test.isErrExpected {
			assert.Error(t, err)
			assert.Contains(t, err.Error(), test.errText)
		} else {
			assert.NoError(t, err)
		}
	}
}

var testsForGetShortURL = []struct {
	num           int
	description   string
	URL           string
	shortURL      string
	isErrExpected bool
	errText       string
	expected      string
}{
	{1, "all fine", "test.com?someparam=value&iammessi=false", "hash", false, "", "hash"},
	{2, "case when there is no data", "", "hash", false, "", ""},
}

func TestTestDBRepo_GetShortURL(t *testing.T) {
	repo := NewInMemoryDBRepo()

	for _, test := range testsForGetShortURL {
		if test.URL != "" {
			_ = repo.Add(test.URL, test.shortURL)
		}
	}
	for _, test := range testsForGetShortURL {
		shortURL, err := repo.GetShortURL(test.URL)
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

var testsForGetLongURL = []struct {
	num           int
	description   string
	shortURL      string
	longURL       string
	isErrExpected bool
	errText       string
	expected      string
}{
	{1, "all fine", "hashed", "test.com/all-fine", false, "", "test.com/all-fine"},
	{2, "there is no long URL", "anotherHashed", "", true, "shortURL not found", ""},
}

func TestInMemoryDBRepo_GetLongURL(t *testing.T) {
	repo := NewInMemoryDBRepo()

	for _, test := range testsForGetLongURL {
		if test.longURL != "" {
			_ = repo.Add(test.longURL, test.shortURL)
		}
	}

	for _, test := range testsForGetLongURL {
		originalURL, err := repo.GetLongURL(test.shortURL)
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
