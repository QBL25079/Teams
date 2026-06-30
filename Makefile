SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command

include .env
export

.PHONY: all env-up env-down env-cleanup migrate-up run

all: env-up migrate-up run

env-up:
	docker compose up -d teams-postgres
	Write-Host "Waiting for PostgreSQL..." -ForegroundColor Yellow
	do { Start-Sleep -Seconds 1; $$ready = docker exec teams-env-postgres pg_isready -U ${POSTGRES_USER} -d ${POSTGRES_DB} 2>$$null; } until ($$LASTEXITCODE -eq 0)
	Write-Host "PostgreSQL is ready." -ForegroundColor Green

env-down:
	docker compose down

env-cleanup:
	docker compose down -v --remove-orphans
	if (Test-Path ".\out\pgdata") { Remove-Item -Recurse -Force ".\out\pgdata" }
	Write-Host "Cleaned." -ForegroundColor Green

migrate-up:
	docker run --rm --network teams_teams-network -v "$(CURDIR)/migrations:/migrations" migrate/migrate:v4.19.1 -path=/migrations -database="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@teams-postgres:5432/${POSTGRES_DB}?sslmode=disable" up

run:
	$$env:POSTGRES_HOST="127.0.0.1"; $$env:POSTGRES_PORT="5432"; $$env:POSTGRES_DATABASE="${POSTGRES_DB}"; $$env:POSTGRES_USER="${POSTGRES_USER}"; $$env:POSTGRES_PASSWORD="${POSTGRES_PASSWORD}"; $$env:POSTGRES_TIMEOUT="${POSTGRES_TIMEOUT}"; go run cmd/teamsapp/main.go
