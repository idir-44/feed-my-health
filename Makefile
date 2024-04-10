start:
	docker compose up --build -d 

stop:
	docker compose rm -s -v -f 


reset:
	docker compose rm -s -v -f 
	docker compose up --build -d 