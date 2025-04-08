.DEFAULT_GOAL := help

.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: lint
lint: ## Run golangci-lint fixing issues
	golangci-lint run --fix

.PHONY: tests
tests: ## Run tests
	go test ./... --tags=unit,integration -coverpkg=./...
