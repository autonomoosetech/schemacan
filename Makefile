.DEFAULT_GOAL := test

gen:
	go generate ./...
.PHONY: gen

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

install: test
	go install -v ./...
.PHONY: install
