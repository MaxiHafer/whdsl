DOCKER_COMPOSE := docker-compose -p whdsl -f deploy/docker-compose.yaml

.PHONY: run
run: all
	bin/api

.PHONY: all
all: swagger build

.PHONY: build
build:
	go build -o bin/api cmd/whdsl/main.go

.PHONY: run
run: all

.PHONY: swagger
swagger:
	go run github.com/swaggo/swag/cmd/swag@latest init --dir cmd/whdsl,$(shell pwd) --output docs --requiredByDefault

.PHONY: compose
compose:
	$(DOCKER_COMPOSE) pull
	$(DOCKER_COMPOSE) up -d
	sleep 3
	$(DOCKER_COMPOSE) logs -f

.PHONY: remove
remove:
	$(DOCKER_COMPOSE) down -v