package service

import (
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/yourusername/urlshortener/internal/model"
	"github.com/yourusername/urlshortener/internal/repository"
)

type URLService struct {
	repo *repository.PostgresRepository
}

func NewURLService(repo *repository.PostgresRepository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) CreateShortURL(longURL string) (*model.URL, error) {
	// Генерация короткого URL
	hash := sha256.Sum256([]byte(longURL + time.Now().String()))
	shortURL := base64.URLEncoding.EncodeToString(hash[:8])

	url := &model.URL{
		LongURL:   longURL,
		ShortURL:  shortURL,
		CreatedAt: time.Now(),
	}

	if err := s.repo.CreateURL(url); err != nil {
		return nil, err
	}

	return url, nil
}

func (s *URLService) GetLongURL(shortURL string) (string, error) {
	url, err := s.repo.GetByShortURL(shortURL)
	if err != nil {
		return "", err
	}

	return url.LongURL, nil
}
