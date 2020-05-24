package corto

import (
	"time"

	"github.com/alwindoss/corto"
)

// NewShortURLManager creates a new short URL Manager
func NewShortURLManager(cfg *corto.Config) corto.ShortURLManager {
	return &shortURLManager{}
}

type shortURLManager struct {
}

func (s *shortURLManager) Create(apiDevKey, originalURL, customAlias, userName string, expireDate time.Time) (string, error) {
	return "", nil
}

func (s *shortURLManager) DeleteURL(apiDevKey, urlKey string) error {
	return nil
}
