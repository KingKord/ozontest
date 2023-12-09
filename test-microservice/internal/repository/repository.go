package repository

type DatabaseRepo interface {
	Add(URL string, shortURL string) error
	GetShortURL(URL string) (string, error)
	GetLongURL(shortURL string) (string, error)
}
