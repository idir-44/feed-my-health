start:
	@docker compose -f docker/docker-compose.yml up --build -d 

stop:
	@docker compose -f docker/docker-compose.yml rm -s -v -f 

reset:
	make stop
	make start 