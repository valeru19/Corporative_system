# Bradobrei Party Backend

Бэкенд информационной системы сети барбершопов `Bradobrei Party` на `Go + Gin + GORM + PostgreSQL/PostGIS`.

## Что уже есть

- REST API для пользователей, сотрудников, салонов, услуг, материалов, бронирований, отзывов, платежей и отчётов.
- JWT-аутентификация с поддержкой `Bearer <token>` (и raw JWT в dev-сценариях).
- Автомиграция таблиц через GORM.
- Swagger UI для локальной разработки.
- PostGIS: `salons.location` как `geometry(Point,4326)`; серверный геокодер для проверки адресов (Yandex и др., см. `.env`).
- HTML/PDF-отчёты через `internal/reports` и Gotenberg.
- E2E и unit-тесты backend-слоя.

## Структура

- `cmd/api/main.go` — точка входа API.
- `cmd/report_example/main.go` — демо рендера HTML/PDF без HTTP.
- `internal/models` — ORM-модели и view-модели для печатных отчётов.
- `internal/dto` — DTO запросов и ответов.
- `internal/handlers` — HTTP-обработчики.
- `internal/services` — прикладная логика.
- `internal/repository` — доступ к данным.
- `internal/geocoder` — провайдеры геокодирования.
- `internal/reports` — шаблоны HTML/CSS и клиент Gotenberg.
- `tests` — e2e/integration тесты.
- `test_artifacts` — сохранённые JSON-артефакты.
- `docs` — сгенерированная Swagger-документация.

## Отчёты 2.2.x из ТЗ

JSON: `GET /api/v1/reports/...`; HTML/PDF: те же пути с суффиксами `/html` и `/pdf`.

- **2.2.1** `GET /employees` — модель печати `models.EmployeeRegistryReportDocument`, шаблон `templates/employees.html`.
- **2.2.2** `GET /salon-activity` — `models.SalonActivityReportDocument`, `salon_activity.html`.
- **2.2.3** `GET /service-popularity` — учитываются бронирования `CONFIRMED`, `IN_PROGRESS`, `COMPLETED`.
- **2.2.4** `GET /master-activity` — `models.MasterActivityReportDocument`, `master_activity.html`.
- **2.2.5** `GET /reviews` — `models.ReviewsReportDocument`, `reviews.html`; создание отзывов: `POST /reviews`.
- **2.2.6** `GET /inventory-movement` — параметры `from`, `to`, `salon_id`.
- **2.2.7** `GET /client-loyalty` — параметры `from`, `to`.
- **2.2.8** `GET /cancelled-bookings` — параметры `from`, `to`.
- **2.2.9** `GET /financial-summary` — параметры `from`, `to`, `salon_id`.

## Зависимости

```bash
go mod download
```

Пересборка Swagger (из каталога `backend`):

```bash
go install github.com/swaggo/swag/cmd/swag@latest
cd cmd/api
swag init -g main.go -o ../../docs -d .,../../internal/handlers,../../internal/dto,../../internal/models
```

## Настройка окружения

Минимальные переменные — в `backend/.env`.

- `DB_*` — PostgreSQL/PostGIS
- `PORT` — порт API (по умолчанию в коде `8080`, если не задан)
- `JWT_SECRET` — подпись JWT
- `GOTENBERG_URL` — Gotenberg для PDF (например `http://localhost:3000`)
- `GEOCODER_PROVIDER`, ключи провайдера — для `POST /api/v1/salons/geocode`

Пример:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=bradobrei
DB_SSLMODE=disable

GIN_MODE=debug
PORT=8080
JWT_SECRET=your-super-secret-jwt-key-change-in-production
GOTENBERG_URL=http://localhost:3000
```

## База данных

```sql
CREATE DATABASE bradobrei;
```

При старте выполняются:

```sql
CREATE SCHEMA IF NOT EXISTS public;
CREATE EXTENSION IF NOT EXISTS postgis;
```

## Локальный запуск

Из папки `backend`:

```bash
go run ./cmd/api
```

Сборка:

```bash
go build -o bradobrei-api ./cmd/api
```

API по умолчанию: `http://localhost:8080` (если `PORT` не переопределён).

## Swagger UI

- `http://localhost:8080/swagger/index.html`
- `http://localhost:8080/docs`

Для защищённых методов: `POST /api/v1/auth/login` → скопировать токен → **Authorize** → `Bearer <jwt>`.

Примеры полей в Swagger задаются в Go-тегах `example` у DTO; после правок перегенерируйте `docs` командой `swag init` выше.

## Тесты

```bash
go test ./internal/...
go test ./tests -v -timeout 60s
go test ./... -v -timeout 120s
```

## Smoke-скрипт отчётов

```bash
./scripts/run_reports_smoke.sh
# или при другом порту:
./scripts/run_reports_smoke.sh http://localhost:8080/api/v1
```

## Gotenberg

```bash
docker compose -f docker-compose.gotenberg.yml up -d
```

Gotenberg обычно на `http://localhost:3000`. В HTML шаблонов ссылайтесь на CSS по имени файла (плоская выгрузка).

## Пример HTML/PDF без HTTP

```bash
go run ./cmd/report_example
go run ./cmd/report_example --skip-pdf
```

## Полезные команды

```bash
gofmt -w ./cmd ./internal ./tests
go build ./...
```

## Частые проблемы

### `schema for creating objects is not selected` / `SQLSTATE 3F000`

- нет схемы `public` или пустой `search_path` у пользователя БД
- неверный `DB_*` в `.env`

### `type "geometry" does not exist`

```sql
CREATE EXTENSION IF NOT EXISTS postgis;
```

### Кеш сборки на Windows

```bash
go clean -cache
```

или:

```powershell
New-Item -ItemType Directory -Force '.gocache' | Out-Null
$env:GOCACHE=(Resolve-Path '.gocache')
go test ./... -v -timeout 120s
```
