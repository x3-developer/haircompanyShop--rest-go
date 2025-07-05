# Hair Company Shop REST API

REST API для магазина бренда Hair Company, разработанное на Go с использованием net/http.

## Описание

Это веб-приложение представляет собой REST API для управления магазином Hair Company. API предоставляет функционал для:

- Клиентской части интернет-магазина
- Административной панели управления

## Технологии

- **Go** 1.24.3
- **GORM** - ORM для работы с базой данных
- **PostgreSQL** - база данных
- **Redis** - кеширование и сессии
- **JWT** - аутентификация
- **Swagger** - документация API
- **Docker** - контейнеризация

## Установка и запуск

### Предварительные требования

- Go 1.24.3 или выше
- PostgreSQL
- Redis

### Шаги установки

1. **Клонирование репозитория**
   ```bash
   git clone <repository-url>
   cd <app-directory>
   ```

2. **Установка зависимостей**
   ```bash
   go mod download
   ```

3. **Настройка окружения**
   ```bash
   cp .env.example .env.local # или .env.production.local - для production
   ```
   Отредактируйте файл `.env*` с вашими настройками.

4. **Запуск миграций**
   ```bash
   go run migrations/auto.go up
   ```
   
5. **Откат миграций (при необходимости)**
   ```bash
   go run migrations/auto.go down
   ```

6. **Запуск приложения**
   ```bash
   go run cmd/main.go
   ```

Приложение будет доступно по адресу: `http://localhost:{APP_PORT}`

## Переменные окружения

Скопируйте `.env.example` в `.env.local` или `.env.production.local` и настройте следующие переменные:

| Переменная | Описание | Обязательная |
|-----------|----------|--------------|
| `APP_ENV` | Окружение приложения (development/production) | ✅ |
| `APP_PORT` | Порт для запуска приложения | ✅ |
| `DB_HOST` | Хост базы данных PostgreSQL | ✅ |
| `DB_PORT` | Порт базы данных PostgreSQL | ✅ |
| `DB_NAME` | Название базы данных | ✅ |
| `DB_USER` | Пользователь базы данных | ✅ |
| `DB_PASSWORD` | Пароль базы данных | ✅ |
| `DB_SSL` | Режим SSL для базы данных | ❌ (по умолчанию: verify-full) |
| `CORS_ALLOWED_ORIGINS` | Разрешенные источники для CORS | ✅ |
| `JWT_DASHBOARD_SECRET_KEY` | Секретный ключ для JWT токенов панели управления | ✅ |
| `JWT_CLIENT_SECRET_KEY` | Секретный ключ для JWT токенов клиентов | ✅ |
| `AUTH_APP_KEY` | Ключ для аутентификации приложения | ✅ |
| `REDIS_ADDR` | Адрес Redis сервера | ✅ |
| `REDIS_PASSWORD` | Пароль Redis | ❌ |
| `REDIS_DB` | Номер базы данных Redis | ❌ (по умолчанию: 0) |

## Структура проекта

```
├── cmd/                    # Входная точка приложения
│   └── main.go
├── config/                 # Конфигурация приложения
│   └── config.go
├── docs/                   # Swagger документация
├── internal/               # Внутренние модули
│   ├── container/          # Dependency injection
│   ├── middleware/         # HTTP middleware
│   ├── modules/            # Бизнес-логика по модулям
│   │   └── v1/
│   │       ├── auth/       # Аутентификация
│   │       ├── category/   # Управление категориями
│   │       ├── client_user/# Пользователи-клиенты
│   │       ├── dashboard_user/ # Пользователи панели управления
│   │       └── image/      # Управление изображениями
│   ├── router/             # HTTP маршруты
│   └── services/           # Общие сервисы
├── migrations/             # Миграции базы данных
├── pkg/                    # Общие пакеты
├── uploads/                # Загруженные файлы
└── go.mod                  # Go модули
```

## API Документация

Swagger документация доступна по адресу: `http://localhost:8080/swagger/index.html`

Для обновления документации выполните:
```bash
cd docs
./generate-swagger.sh
```

## Разработка

### Генерация Swagger документации
```bash
cd docs
./generate-swagger.sh
```

### Форматирование Swagger документации
```bash
  swag fmt
```

### Создание новой миграции
```bash
migrate create -ext sql -dir migrations -seq your_migration_name
```

## Контакты

Поддержка API - x3.na.tri@gmail.com

Ссылка на проект: [https://github.com/x3-developer/haircompanyShop--rest-go](https://github.com/your-username/haircompany-shop-rest)
