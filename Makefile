# Commands
GO      ?= go
GOLINT  ?= golint

.PHONY: default lint vet test

default: lint vet test

lint:
	$(GOLINT) ./...

vet:
	$(GO) vet ./...

test:
	$(GO) test ./...
