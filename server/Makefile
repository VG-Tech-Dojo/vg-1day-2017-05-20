GOOS = darwin
GOARCH = amd64
VERSION = $(shell git rev-parse HEAD)

all: deps

deps:
	which dep || go get -u github.com/golang/dep/...

install:
	dep ensure
	go install    # cache compiled packages

build:
	mkdir -p builds
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o builds/server -ldflags "-X=main.version=$(VERSION)" server.go

clean: clean_vendor

clean_vendor:
	rm -rf vendor/
