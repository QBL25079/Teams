# Teams REST API

**Автор:** Шульгин Николай

REST API сервис для управления группами людей.

## Описание проекта

Сервис реализует систему управления группами людей.

Основные возможности:

- создание, получение, обновление и удаление групп;
- поддержка вложенных (дочерних) групп;
- создание, получение, обновление и удаление пользователей;
- привязка пользователя к группе;
- изменение группы пользователя;
- получение списка групп;
- получение списка пользователей;
- получение пользователей конкретной группы;
- получение пользователей группы вместе со всеми дочерними группами;
- подсчёт количества пользователей как непосредственно в группе, так и с учётом всех дочерних групп.

Проект реализован в соответствии с принципами многослойной архитектуры:

- Transport (HTTP)
- Service
- Repository
- Domain

---

# Используемые технологии

- Go 1.26.1
- PostgreSQL 17
- Docker
- Docker Compose
- golang-migrate
- pgx
- zap

---

# Структура проекта

```text
cmd/
    teamsapp/
internal/
    core/
        domain/
        errors/
        logger/
        repository/
            transport/
                postgres/
                    pool/
                        pgx/
            http/
                middleware/
                request/
                response/
                server/
                types/
    features/
        teams/
            repository/
            service/
            transport/
                http/
        users/
            repository/
            service/
            transport/
                http/
migrations/
go.mod
go.sum
docker-compose.yml
Makefile
README.md
```

---

# Подготовительные действия

Для запуска необходимо установить:

- Go 1.26+
- Docker
- Docker Compose

---

# Настройка проекта

Создать базу данных не требуется.

При запуске Docker Compose PostgreSQL будет создан автоматически.

После запуска контейнеров необходимо применить миграции.

---

# Доступы

## PostgreSQL

| Параметр | Значение |
|----------|----------|
| Host | localhost |
| Port | 5433 |
| Database | test-db |
| User | test-user-123 |
| Password | pass |

---

# Запуск проекта

## 1. Запустить PostgreSQL

```bash
docker compose up -d
```

или

```bash
docker compose up --build -d
```

---

## 2. Применить миграции

Если используется Makefile:

```bash
make migrate-up
```

или напрямую:

```bash
migrate \
-path ./migrations \
-database "postgres://test-user-123:pass@localhost:5433/test-db?sslmode=disable" \
up
```

---

## 3. Запустить приложение

Если используется Makefile:

```bash
make teams-run
```

или

```bash
go run cmd/teamsapp/main.go
```

---

После запуска сервис будет доступен по адресу

```
http://localhost:8080 или 127.0.0.1:8080
```

---

# REST API

## Команды

| Метод | Endpoint | Описание |
|--------|----------|----------|
| POST | `/api/v1/team` | Создать команду |
| GET | `/api/v1/teams` | Получить список команд |
| GET | `/api/v1/team/{id}` | Получить команду |
| PATCH | `/api/v1/team/{id}` | Обновить команду |
| DELETE | `/api/v1/team/{id}` | Удалить команду |

---

## Пользователи

| Метод | Endpoint | Описание |
|--------|----------|----------|
| POST | `/api/v1/user` | Создать пользователя |
| GET | `/api/v1/users` | Получить список пользователей |
| GET | `/api/v1/user/{id}` | Получить пользователя |
| PATCH | `/api/v1/user/{id}` | Обновить пользователя |
| DELETE | `/api/v1/user/{id}` | Удалить пользователя |

---

# Возможности

Проект поддерживает:

- создание дочерних команд;
- изменение родительской команды;
- удаление команды;
- автоматическое обнуление `parent_id` дочерних команд после удаления родительской (через `ON DELETE SET NULL`);
- привязку пользователя к команде;
- изменение команды пользователя;
- получение пользователей команды;
- получение пользователей вместе с пользователями всех дочерних команд;
- подсчёт количества пользователей.

---

# Миграции

Применить

```bash
make migrate-up
```

или

```bash
migrate -path ./migrations \
-database "postgres://test-user-123:pass@localhost:5433/test-db?sslmode=disable" \
up
```

Откатить

```bash
make migrate-down
```

или

```bash
migrate -path ./migrations \
-database "postgres://test-user-123:pass@localhost:5433/test-db?sslmode=disable" \
down
```

---

# Примечания

- Все ответы возвращаются в формате JSON.
- Используется PostgreSQL.
- Для управления схемой базы данных используется **golang-migrate**.
- Все временные значения хранятся в формате UTC (`TIMESTAMPTZ`).

---