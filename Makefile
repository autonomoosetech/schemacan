.DEFAULT_GOAL := help

help: ## Self documenting help output
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.PHONY: help

gen: ## Generate code from OpenAPI spec
	go generate ./...
.PHONY: gen

fmt: ## Format code
	go fmt ./...
.PHONY: fmt

lint: fmt ## Lint code
	golint ./...
.PHONY: lint

vet: fmt ## Vet code
	go vet ./...
.PHONY: vet

test: vet ## Run tests
	go test -v ./...
.PHONY: test

install: test ## Install as executable
	go install -v ./...
.PHONY: install
