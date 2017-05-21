.DEFAULT_GOAL := help

GOOS = darwin
GOARCH = amd64
VERSION = $(shell git rev-parse HEAD)
ENV = test
host=localhost:8080

server:
	go run server.go

## Live reload
watch:
	fresh

## Install dependency tools
deps:
	which make2help || go get -u github.com/Songmu/make2help/cmd/make2help
	which dep || go get -u github.com/golang/dep/...
	which sql-migrate || go get -u github.com/rubenv/sql-migrate/...
	which fresh || go get -u github.com/pilu/fresh
	which golint || go get -u github.com/golang/lint/golint

## Install dependency packages, cache compiled packages
install: deps env/env.go dev.db 
	dep ensure
	go install

## Build app
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X=main.version=$(VERSION)" server.go

## Check code format
check: fmt vet lint

fmt:
	go fmt $$(go list ./... | grep -v vendor)

vet:
	go vet $$(go list ./... | grep -v vendor)

lint:
	golint $$(go list ./... | grep -v vendor)

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
	curl -i -X POST $(host)/api/messages -d '{"body": "$(body)","sender_name": "$(sender_name)"}'

curl_message_put:
	curl -i -X PUT $(host)/api/messages/$(id) -d '{"body": "$(body)","sender_name": "$(sender_name)"}'

curl_message_delete:
	curl -i -X DELETE $(host)/api/messages/$(id) -d '{"body": "$(body)"}'

dev.db:
	$(MAKE) db_init

## Initialize database
db_init:
	cp -i _etc/seed.db dev.db

## Initialize environment configuration file
env/env.go:
	cp env/env.go.tmpl env/env.go

## Build godoc server(http://localhost:6060)
godoc_server:
	godoc -http=:6060
