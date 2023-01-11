DOCKER_COMPOSE_CMD := docker-compose \
	-f docker-compose.yml \
	-f docker-compose.local.yml
DOCKER_COMPOSE_TSS_CMD := docker-compose \
	-f docker-compose.tss.yml
DOCKER_COMPOSE_L1_CMD := docker-compose \
	-f docker-compose.l1.yml
DOCKER_COMPOSE_L2_CMD := docker-compose \
	-f docker-compose.l2.yml
COPY_ENV := cp -r envs/deployer.env ./.env
COPY_CHAIN := cp -r chaintestdata ./chaindata

build:
	DOCKER_BUILDKIT=1 \
	$(DOCKER_COMPOSE_CMD) \
		build $(service)
.PHONY: build

up:
	$(COPY_ENV)
	$(CLEAR_DOCKER_NETWORK)
	$(COPY_CHAIN)
	DOCKER_BUILDKIT=1 \
	$(DOCKER_COMPOSE_CMD) \
		up --build --detach $(service)
.PHONY: up


l1up:
	$(COPY_ENV)
	$(COPY_CHAIN)
	DOCKER_BUILDKIT=1 \
	$(DOCKER_COMPOSE_L1_CMD) \
		up --build --detach $(service)
.PHONY: up


l2up:
	DOCKER_BUILDKIT=1 \
	$(DOCKER_COMPOSE_L2_CMD) \
		up --build --detach $(service)
.PHONY: up


up-tss:
	DOCKER_BUILDKIT=1 \
	$(DOCKER_COMPOSE_TSS_CMD) \
		up --build --detach $(service)
.PHONY: up-tss


down:
	$(DOCKER_COMPOSE_CMD) \
		down
.PHONY: down


l2down:
	$(DOCKER_COMPOSE_L2_CMD) \
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
	docker network rm ops_mantle
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


