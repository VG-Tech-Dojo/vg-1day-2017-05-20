SED=$(shell which sed)
.DEFAULT_GOAL := help
background_option=-d
nickname=
repository_name=$(shell basename $(PWD))

setup/mac:
	$(MAKE) setup/bsd

setup/bsd: $(nickname) ## for mac
	$(SED) -i '' -e 's/original/$(nickname)/g' ./$(nickname)/*.go
	$(SED) -i '' -e 's/original/$(nickname)/g' ./$(nickname)/**/*.go
	$(SED) -i '' -e 's/vg-1day-2017/$(repository_name)/g' ./$(nickname)/*.go
	$(SED) -i '' -e 's/vg-1day-2017/$(repository_name)/g' ./$(nickname)/**/*.go

setup/gnu: $(nickname) ## for linux
	$(SED) --in-place 's/original/$(nickname)/g' ./$(nickname)/*.go
	$(SED) --in-place 's/original/$(nickname)/g' ./$(nickname)/**/*.go
	$(SED) --in-place 's/vg-1day-2017/$(repository_name)/g' ./$(nickname)/*.go
	$(SED) --in-place 's/vg-1day-2017/$(repository_name)/g' ./$(nickname)/**/*.go

$(nickname):
	cp -rf original $(nickname)

docker_server: docker_build docker_up

docker_clean: docker_stop docker_rm 

help:
	@echo docker_build:	Build the docker container
	@echo docker_up:	Start the docker container
	@echo docker_stop:	Stop the docker container
	@echo docker_rm:	Remove the docker container
	@echo docker_ssh:	Execute an interactive bash shell on the container

docker_build:
	docker-compose build

docker_up:
	docker-compose up $(background_option)

docker_stop:
	docker-compose stop

docker_rm:
	docker-compose rm

docker_ssh:
	docker exec -it vg-1day-2017-go /bin/bash
