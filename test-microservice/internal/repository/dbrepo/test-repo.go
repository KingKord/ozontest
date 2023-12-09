package dbrepo

import "errors"

func (t testDBRepo) Add(URL string, shortURL string) error {
	if URL == "test.com/duplicate-hash" {
		return errors.New("duplicate hashes")
	}

	return nil
}

func (t testDBRepo) GetShortURL(URL string) (string, error) {
	if URL == "test.com/existing" {
		return "existingHash", nil
	}

	if URL == "test.com/err-in-DB" {
		return "", errors.New("error in DB")
	}

	return "", nil
}

func (t testDBRepo) GetLongURL(shortURL string) (string, error) {
	if shortURL == "non-existenceHash" {
		return "", errors.New("hash not exist in DB")
	}

	if shortURL == "invalidDB" {
		return "", errors.New("invalid DB")
	}

	return "youtube.com/channels/beast?redirect=true&from=finder", nil
}
