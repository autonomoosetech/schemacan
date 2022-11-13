.DEFAULT_GOAL := help

.PHONY: help
help: ## Self documenting help output
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: gen
gen: ## Generate code from OpenAPI spec
	go generate ./...
	sed -i '' "s/json:/yaml:/" api/v1/api.gen.go # this is to support yaml files

.PHONY: fmt
fmt: ## Format code
	go fmt ./...

.PHONY: lint
lint: fmt ## Lint code
	golint ./...

.PHONY: vet
vet: fmt ## Vet code
	go vet ./...

.PHONY: test
test: vet ## Run tests
	go test -v ./...

.PHONY: install
install: gen test ## Install as executable
	go install -v ./...

.PHONY: build
build: gen test ## Build executable
	go build -v -o bin/ ./...
