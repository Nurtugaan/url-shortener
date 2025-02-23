package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/urlshortener/config"
	"github.com/yourusername/urlshortener/internal/handler"
	"github.com/yourusername/urlshortener/internal/repository"
	"github.com/yourusername/urlshortener/internal/service"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Инициализация репозитория
	repo, err := repository.NewPostgresRepository(cfg.DB)
	if err != nil {
		log.Fatalf("Error creating repository: %v", err)
	}

	// Инициализация сервиса
	svc := service.NewURLService(repo)

	// Инициализация обработчика
	h := handler.NewHandler(svc)

	// Настройка роутера
	r := gin.Default()

	// Роуты
	r.POST("/shorten", h.CreateShortURL)
	r.GET("/:shortURL", h.RedirectToLongURL)

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
