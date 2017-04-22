.DEFAULT_GOAL := help

GOOS = darwin
GOARCH = amd64
VERSION = $(shell git rev-parse HEAD)
ENV = test
host=localhost:8080

server: install
	go run server.go

## Install dependency tools
deps:
	which make2help || go get -u github.com/Songmu/make2help/cmd/make2help
	which dep || go get -u github.com/golang/dep/...
	which sql-migrate || go get -u github.com/rubenv/sql-migrate/...
	which fresh || go get -u github.com/pilu/fresh

## Install dependency packages, cache compiled packages
install: deps
	dep ensure
	go install

## Build app
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X=main.version=$(VERSION)" server.go

## Live reload
watch:
	fresh

clean: clean_vendor

## Clean vendor directory
clean_vendor:
	rm -rf vendor/

## Show help
help:
	@make2help -all

## Migrate db schema
migrate_up:
	sql-migrate up -env=$(ENV)

## Migrate db schema(dryrun)
migrate_dryrun:
	sql-migrate up -env=$(ENV) -dryrun

## Show migration status
migrate_status:
	sql-migrate status -env=$(ENV)

curl_messages_ping:
	curl -i $(host)/api/ping

curl_messages_get_all:
	curl -i $(host)/api/messages

curl_messages_get:
	curl -i $(host)/api/messages/$(id)

curl_message_create:
	curl -i -X POST $(host)/api/messages -d '{"body": "$(body)"}'

