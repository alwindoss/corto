package corto

import "time"

// Config is the corto Config
type Config struct {
}

// ShortURLManager interface defines
type ShortURLManager interface {
	Create(apiDevKey, originalURL, customAlias, userName string, expireDate time.Time) (string, error)
	DeleteURL(apiDevKey, urlKey string) error
}
