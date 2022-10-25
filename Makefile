DOCKER_COMPOSE := docker compose -p whdsl -f deploy/docker-compose.yaml

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

.PHONY: client
client:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	oapi-codegen -generate client,types -o ./pkg/client/client_gen.go -package client --old-config-style ./openapi.yaml


.PHONY: compose
compose:
	$(DOCKER_COMPOSE) pull
	$(DOCKER_COMPOSE) up -d
	sleep 3
	$(DOCKER_COMPOSE) logs -f

.PHONY: remove
remove:
	$(DOCKER_COMPOSE) down -v