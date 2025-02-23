# URL Shortener

URL Shortener - это сервис для сокращения длинных URL-адресов с использованием Golang, PostgreSQL и Gin.

## Возможности
- Сокращение длинных URL.
- Перенаправление по короткому URL.
- Сохранение данных в базе PostgreSQL.
- Запуск с помощью Docker и Docker Compose.

## Структура проекта
```
url-shortener/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── config/
│   └── config.go
├── internal/
│   ├── handler/
│   │   └── handler.go
│   ├── model/
│   │   └── url.go
│   ├── repository/
│   │   └── postgres.go
│   └── service/
│       └── service.go
└── migrations/
    └── 001_create_urls_table.sql
```

## Установка и запуск
### 1. Клонирование репозитория
```sh
git clone https://github.com/yourusername/urlshortener.git
cd urlshortener
```

### 2. Запуск с помощью Docker Compose
```sh
docker-compose up --build
```

Приложение будет доступно по адресу `http://localhost:8080`.

### 3. Запуск без Docker
#### Установка зависимостей
```sh
go mod tidy
```
#### Запуск миграций
Перед запуском сервера необходимо создать таблицы в базе данных. Подключитесь к PostgreSQL и выполните SQL-скрипт:
```sh
psql -h localhost -U postgres -d urlshortener -f migrations/001_create_urls_table.sql
```

#### Запуск сервера
```sh
go run main.go
```

## Использование API

### 1. Создание короткого URL
**POST /shorten**

#### Запрос:
```json
{
  "long_url": "https://example.com/some/very/long/url"
}
```

#### Ответ:
```json
{
  "id": 1,
  "long_url": "https://example.com/some/very/long/url",
  "short_url": "abc123",
  "created_at": "2025-02-23T12:34:56Z"
}
```

### 2. Перенаправление по короткому URL
**GET /:shortURL**

#### Пример запроса:
```sh
curl -L http://localhost:8080/abc123
```

Перенаправит пользователя на `https://example.com/some/very/long/url`.

## Конфигурация
Файл `config/config.go` загружает переменные окружения:
- `DB_HOST` - хост базы данных
- `DB_PORT` - порт базы данных
- `DB_USER` - пользователь базы данных
- `DB_PASSWORD` - пароль базы данных
- `DB_NAME` - имя базы данных

Переменные можно настроить в `docker-compose.yml` или передавать через `.env`.

## Технологии
- **Golang** - язык программирования
- **Gin** - веб-фреймворк
- **PostgreSQL** - база данных
- **Docker** и **Docker Compose** - контейнеризация

