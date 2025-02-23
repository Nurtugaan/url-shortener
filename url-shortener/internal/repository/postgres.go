package repository

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/yourusername/urlshortener/config"
	"github.com/yourusername/urlshortener/internal/model"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(cfg config.DatabaseConfig) (*PostgresRepository, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	// Добавляем retry логику
	var db *sql.DB
	var err error
	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			time.Sleep(time.Second * 2)
			continue
		}

		if err = db.Ping(); err != nil {
			time.Sleep(time.Second * 2)
			continue
		}

		return &PostgresRepository{db: db}, nil
	}

	return nil, fmt.Errorf("failed to connect to database after 5 attempts: %v", err)
}

func (r *PostgresRepository) CreateURL(url *model.URL) error {
	query := `
        INSERT INTO urls (long_url, short_url, created_at)
        VALUES ($1, $2, $3)
        RETURNING id`

	return r.db.QueryRow(query, url.LongURL, url.ShortURL, url.CreatedAt).Scan(&url.ID)
}

func (r *PostgresRepository) GetByShortURL(shortURL string) (*model.URL, error) {
	url := &model.URL{}
	query := `
        SELECT id, long_url, short_url, created_at
        FROM urls
        WHERE short_url = $1`

	err := r.db.QueryRow(query, shortURL).Scan(&url.ID, &url.LongURL, &url.ShortURL, &url.CreatedAt)
	if err != nil {
		return nil, err
	}

	return url, nil
}
