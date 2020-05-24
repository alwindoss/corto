package corto

import (
	"log"
	"sync"
	"time"

	"github.com/alwindoss/corto"
	"github.com/teris-io/shortid"
)

var singleton sync.Once

// NewShortURLManager creates a new short URL Manager
func NewShortURLManager(cfg *corto.Config) corto.ShortURLManager {
	var mgr corto.ShortURLManager
	var e error
	singleton.Do(func() {
		urlDB := make(map[string]*urlDetails)
		userDB := make(map[string]string)
		gen, err := shortid.New(1, shortid.DefaultABC, 2342)
		if err != nil {
			e = err
			return
		}
		mgr = &shortURLManager{
			generator: gen,
			urlDB:     urlDB,
			userDB:    userDB,
		}
	})
	if e != nil {
		return nil
	}
	return mgr
}

type urlDetails struct {
	expiryDate  time.Time
	userName    string
	shortURLKey string
	originalURL string
	customAlias string
	apiDevKey   string
}

type shortURLManager struct {
	sync.RWMutex
	generator *shortid.Shortid
	urlDB     map[string]*urlDetails
	userDB    map[string]string
}

func (s *shortURLManager) Create(apiDevKey, originalURL, customAlias, userName string, expiryDate time.Time) (string, error) {
	key, err := s.generator.Generate()
	if err != nil {
		log.Printf("unable to generate a shortid")
		return "", err
	}
	ud := &urlDetails{
		apiDevKey:   apiDevKey,
		originalURL: originalURL,
		shortURLKey: key,
		customAlias: customAlias,
		userName:    userName,
		expiryDate:  expiryDate,
	}
	s.Lock()
	s.urlDB[key] = ud
	defer s.Unlock()
	return key, nil
}

func (s *shortURLManager) DeleteURL(apiDevKey, urlKey string) error {
	s.Lock()
	delete(s.urlDB, urlKey)
	defer s.Unlock()
	return nil
}
func (s *shortURLManager) FetchURL(apiDevKey, urlKey string) (string, error) {
	var ud *urlDetails
	s.Lock()
	ud = s.urlDB[urlKey]
	originalURL := ud.originalURL
	defer s.Unlock()
	return originalURL, nil
}
