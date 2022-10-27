.PHONY: run
run: build
	./bin/whdsl

.PHONY: build
build:
	go build -o bin/whdsl cmd/whdsl/main.go

.PHONY: generate
generate:
	go generate ./...

.PHONY: swagger-editor
swagger-editor:
	docker run --rm -p 8081:8080 -e URL="http://localhost:8080/swagger/openapi.yaml" swaggerapi/swagger-editor
