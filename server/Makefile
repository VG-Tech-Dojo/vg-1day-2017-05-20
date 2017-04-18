.DEFAULT_GOAL := help

GOOS = darwin
GOARCH = amd64
VERSION = $(shell git rev-parse HEAD)

server: install
	go run server.go


## Install dependency tools
deps:
	which make2help || go get -u github.com/Songmu/make2help/cmd/make2help
	which dep || go get -u github.com/golang/dep/...

## Install dependency packages, cache compiled packages
install: deps
	dep ensure
	go install

## Build app
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X=main.version=$(VERSION)" server.go

clean: clean_vendor

## Clean vendor directory
clean_vendor:
	rm -rf vendor/

## Show help
help:
	@make2help -all
