SHELL=/bin/bash

all: build

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: run
run: tidy
	go run main.go $(MOD)

.PHONY: build
build: tidy
	go build

.PHONY: test
test: tidy
	pushd source && go test