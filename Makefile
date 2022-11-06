.PHONY: run
run: build
	./bin/whdsl

.PHONY: dep
dep:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@master

.PHONY: build
build:
	go build -o bin/whdsl cmd/whdsl/main.go

.PHONY: buf
buf: buf-push

.PHONY: buf-lint
buf-lint:
	docker run -v $(shell pwd):/work --workdir /work bufbuild/buf lint proto

.PHONY: buf-generate
buf-generate: buf-lint
	docker run -v $(shell pwd):/work --workdir /work/proto bufbuild/buf generate
	sudo chown -R $(shell id -u):$(shell id -g) pkg/pb

.PHONY: buf-push
buf-push: buf-lint buf-generate
	docker run -v $(HOME)/.netrc:/root/.netrc -v $(shell pwd):/work --workdir /work bufbuild/buf push proto

.PHONY: generate
generate:
	go generate ./...