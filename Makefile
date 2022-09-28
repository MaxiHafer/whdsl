DOCKER_COMPOSE := docker compose -p whdsl -f deploy/docker-compose.yaml -f deploy/docker-compose.override.yaml

.SPHONY: swagger
swagger:
	swag init

.PHONY: compose
compose:
	$(DOCKER_COMPOSE) pull
	$(DOCKER_COMPOSE) up -d
	sleep 3
	$(DOCKER_COMPOSE) logs -f

.PHONY: remove
remove:
	$(DOCKER_COMPOSE) down -v