package service

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"strings"
	"time"
	"url-shortener/internal/repository"
)

var (
	ErrEmptyURL = errors.New("empty url")
)

type URLService struct {
	repo repository.URLRepository
}

func NewURLService(r repository.URLRepository) *URLService {
	return &URLService{repo: r}
}

func (s *URLService) CreateShortURL(ctx context.Context, originalURL, alias string) (string, error) {
	originalURL = strings.TrimSpace(originalURL)
	if originalURL == "" {
		return "", ErrEmptyURL
	}

	if alias == "" {
		alias = generateAlias(originalURL)
	}

	if err := s.repo.Create(ctx, originalURL, alias); err != nil {
		return "", err
	}
	return alias, nil
}

func (s *URLService) GetOriginalURL(ctx context.Context, alias string) (string, error) {
	return s.repo.Get(ctx, alias)
}

func (s *URLService) DeleteShortURL(ctx context.Context, alias string) error {
	return s.repo.Delete(ctx, alias)
}

func generateAlias(u string) string {
	h := sha1.New()
	h.Write([]byte(u + time.Now().String()))
	return hex.EncodeToString(h.Sum(nil))[:7]
}
