EXEC ?=docker compose -f docker/docker-compose.yml exec -e CGO_ENABLED=0 -T docker-server-1

start:
	@docker compose -f docker/docker-compose.yml up --build -d 

stop:
	@docker compose -f docker/docker-compose.yml rm -s -v -f 

migration-init:
	docker exec -it docker-server-1 go run ./cmd/migrate/main.go db init

migration-up:
	docker exec -it docker-server-1 go run ./cmd/migrate/main.go db migrate

reset:
	make stop
	make start 