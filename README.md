# Сервис для обработки заказов

## Запуск

Запустите приложение:

```bash
make run
```

Запустите скрипт для публикации сообщений в канал:

```bash
make run-publisher
```

→ [localhost:1337](http://localhost:1337)

## Локальная разработка

Установите необходимые инструменты: `air`, `gofumpt`, `golangci-lint`, `pre-commit`:

```bash
make init
```

Запустите приложение в live-режиме:

```bash
make run-air
```

Дополнительно:

|   |   |
|---|---|
| `make deps` | Обновить зависимости |
| `make lint` | Запустить линтеры для всего проекта |
| `make migrate-up` | Применить миграции |
| `make migrate-down` | Откатить миграции |
| `make` | Показать `help` |

## Стек технологий

- Go `1.21`
- Python `3.11`
- Docker
- `make` в качестве таск-раннера

Пакеты:

- [`echo`](https://github.com/labstack/echo) — легковесный веб-фреймворк
- [`godotenv`](https://github.com/joho/godotenv) — для конфигурации приложения через переменные окружения, чтоби приблизить его к 12-факторному
- [`stan.go`](github.com/nats-io/stan.go) — подключение к nats-streaming
- [`pq`](https://github.com/lib/pq) — драйвер для `PostgreSQL`
- [`log/slog`](https://pkg.go.dev/log/slog) — для структурированного логгирования
- [`console-slog`](https://github.com/phsym/console-slog) — форматтер для логов для локальной разработки
- [`slog-echo`](https://github.com/samber/slog-echo) — мидлварь для `echo`, объединяющий логи в `slog`
- [`validator`](github.com/go-playground/validator/v10) — валидатор сообщений из канала

Инструменты:

- [`air`](https://github.com/cosmtrek/air) — автоматическая пересборка приложения
- [`gofumpt`](https://github.com/mvdan/gofumpt) — форматтер, построже, чем `gofmt`
- [`golangci-lint`](https://github.com/golangci/golangci-lint) — металинтер для go
- [`pre-commit`](https://github.com/pre-commit/pre-commit) — хуки для валидации кода перед каждый коммитом
