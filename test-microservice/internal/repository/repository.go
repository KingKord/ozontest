package repository

type DatabaseRepo interface {
	Add(URL string, shortURL string) error
	GetShortURL() (string, error)
	GetLongURL(shortURL string) (string, error)
}
