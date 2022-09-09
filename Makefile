.DEFAULT_GOAL := test

fmt:
	go fmt ./...
.PHONY: fmt

lint: fmt
	golint ./...
.PHONY: lint

vet: fmt
	go vet ./...
.PHONY: vet

test: vet
	go test -v ./...
.PHONY: test
