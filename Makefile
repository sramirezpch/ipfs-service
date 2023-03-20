IMAGE_NAME=ipfs-service
DOCKERFILE_PATH=./infra/.dockerfile
COMPOSE_FILE=./infra/docker-compose.yml
CONTAINER_NAME=ipfs-service

build-image:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE_PATH). --no-cache

run-with-docker:
	make build-image
	docker run --name $(ipfs-service) -p 3002:8080 --rm -it -d $(IMAGE_NAME)

down:
	docker-compose down --remove-orphans

.PHONY: run
run:
	make down
	docker-compose up -f $(COMPOSE_FILE) up -d --build