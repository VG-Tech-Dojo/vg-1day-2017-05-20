.DEFAULT_GOAL := help


help:
	@echo docker_build:	Build the docker container
	@echo docker_up:	Start the docker container
	@echo docker_stop:	Stop the docker container
	@echo docker_rm:	Remove the docker container
	@echo docker_ssh:	Execute an interactive bash shell on the container

docker_build:
	docker-compose build

docker_up:
	docker-compose up -d

docker_stop:
	docker-compose stop

docker_rm:
	docker-compose rm

docker_ssh:
	docker exec -it vg-1day-2017-go /bin/bash
