.PHONY: help build tests gen

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## build goval
	go build -o ./dist ./cmd/goval

gen: ## generate
	go generate ./tests/...

tests: build ## build and run tests
	go generate ./tests/...
	go test -v ./...
