.PHONY: start-db
start-db:
	docker-compose up

.PHONY: drop-db
drop-db:
	docker-compose down

.PHONY: runserver
runserver:
	go run .\cmd\main.go