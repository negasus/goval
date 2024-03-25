.PHONY: help build tests

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build:
	go build -o ./dist ./cmd/goval

tests: build
	go generate ./tests/...
	go test -v ./...
