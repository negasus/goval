.PHONY: help example build

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

example: build ## Build and run an example
	go generate ./...

build:
	go build -o ./dist ./cmd/goval
