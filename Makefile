PARAMS := $(wordlist 2,100,${MAKECMDGOALS})

.PHONY: start-db
start-db:
	docker-compose up -d migrate ${PARAMS}

.PHONY: drop-db
drop-db:
	docker-compose rm --stop --force

.PHONY: start-app
start-app:
	docker-compose up -d app ${PARAMS}

.PHONY: runserver
runserver:
	go run ./cmd/server/main.go -env="dev"

.PHONY: build
build:
	go build ./cmd/server/main.go