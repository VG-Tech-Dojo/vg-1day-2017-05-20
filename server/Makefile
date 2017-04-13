GLIDE = $(shell which glide)

all: deps

deps:
	which glide || curl -sS https://glide.sh/get | sh
	$(GLIDE) install
