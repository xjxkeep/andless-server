package main

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

// CaptchaStore implements base64Captcha.Store using go-cache.
type CaptchaStore struct {
	cache *gocache.Cache
}

func NewCaptchaStore() *CaptchaStore {
	return &CaptchaStore{
		cache: gocache.New(5*time.Minute, 10*time.Minute),
	}
}

// Set stores the captcha answer with the given id.
func (s *CaptchaStore) Set(id string, value string) error {
	s.cache.Set(id, value, gocache.DefaultExpiration)
	return nil
}

// Get retrieves the captcha answer and optionally removes it.
func (s *CaptchaStore) Get(id string, clear bool) string {
	val, found := s.cache.Get(id)
	if !found {
		return ""
	}
	if clear {
		s.cache.Delete(id)
	}
	return val.(string)
}

// Verify checks the answer and removes the captcha entry (prevents replay).
func (s *CaptchaStore) Verify(id, answer string, clear bool) bool {
	stored := s.Get(id, clear)
	return stored != "" && stored == answer
}
