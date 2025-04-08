.DEFAULT_GOAL := help

.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: lint
lint: ## Run golangci-lint fixing issues
	golangci-lint run --fix

.PHONY: tests
tests: ## Run tests checking race conditions
	go test --tags=unit,integration -coverpkg=$(go list ./... | grep -v /example/)

.PHONY: example
example: ## Run example
	go run ./example/main.go

.PHONY: racetests
racetests: ## Run tests with race condition checking
	CGO_ENABLED=1 go test -race --tags=unit,integration -coverpkg=$(go list ./... | grep -v /example/)