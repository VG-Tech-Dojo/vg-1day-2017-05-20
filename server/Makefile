GLIDE = $(shell which glide)

all: deps

deps:
	which dep || go get -u github.com/golang/dep/...

install:
	dep ensure
	go install    # cache compiled packages

clean: clean_vendor

clean_vendor:
	rm -rf vendor/
