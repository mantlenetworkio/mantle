DOCKER_COMPOSE_CMD := docker-compose \
	-f docker-compose.yml \
	-f docker-compose.local.yml
COPY_ENV := cp -r envs/deployer.env ./.env
COPY_CHAIN := cp -r chaintestdata ./chaindata

build:
	DOCKER_BUILDKIT=1 \
	$(DOCKER_COMPOSE_CMD) \
		build $(service)
.PHONY: build

up:
	$(COPY_ENV)
	$(COPY_CHAIN)
	DOCKER_BUILDKIT=1 \
	$(DOCKER_COMPOSE_CMD) \
		up --build --detach $(service)
.PHONY: up

down:
	$(DOCKER_COMPOSE_CMD) \
		down
.PHONY: down

ps:
	$(DOCKER_COMPOSE_CMD) \
		ps
.PHONY: ps

start:
	$(DOCKER_COMPOSE_CMD) \
		start $(service)
.PHONY: restart

stop:
	$(DOCKER_COMPOSE_CMD) \
		stop $(service)
.PHONY: restart

restart:
	$(DOCKER_COMPOSE_CMD) \
		restart $(service)
.PHONY: restart

clean: stop
	sudo rm -rf chaindata
	find data -mindepth 1 -maxdepth 1 -type d | xargs rm -rf
.PHONY: clean

up-metrics: down-metrics
	DOCKER_BUILDKIT=1 \
	$(DOCKER_COMPOSE_CMD) \
		-f docker-compose-metrics.yml \
		up --build --detach
.PHONY: up-metrics

down-metrics:
	$(DOCKER_COMPOSE_CMD) \
		-f docker-compose-metrics.yml \
		down
.PHONY: down-metrics

ps-metrics:
	$(DOCKER_COMPOSE_CMD) \
		-f docker-compose-metrics.yml \
		ps
.PHONY: ps


