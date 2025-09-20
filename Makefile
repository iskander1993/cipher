.PHONY: run up down logs build restart status

run:
	go run cmd/main.go
#Поднимает контейнеры
up:
	docker-compose up -d

#Останавливаем контейнеры
down:
	docker-compose down

#Логи контейнеров
logs:
	docker-compose logs -f

logs-db:
	docker-compose logs -f db

build:
	docker-compose up -d --build

restart:
	docker-compose restart

status:
	docker-compose ps