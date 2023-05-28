.PHONY: all
all: deps lint test

.PHONY: run
run:
	go run .

.PHONY: run-dev
run-dev:
	command -v wgo || go install github.com/bokwoon95/wgo@latest
	wgo run .

.PHONY: generate
generate:
	go generate ./...

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

.PHONY: run-redis
run-redis: ## Run redis and RedisInsight
	docker run -d --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
