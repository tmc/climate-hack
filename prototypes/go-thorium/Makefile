.PHONY: all
all: deps lint test

.PHONY: run
run:
	go run .

.PHONY: run-dev
run-dev:
	command -v wgo || go install github.com/bokwoon95/wgo@latest
	wgo run .

.PHONY: deps
deps:
	command -v go || brew install go
	command -v golangci-lint || brew install golangci-lint

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -v ./...
