.PHONY: default run build test docs stack clean

#Variables
APP_NAME=ps-tag-onboarding-go
COMPOSE_FILE=.docker/stack.yaml

# Tasks
default: run-with-docs

run:
	@docker compose -f $(COMPOSE_FILE) up -d
run-with-docs:
	@swag init
	@docker compose -f $(COMPOSE_FILE) up -d
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
stack:
	@docker compose -f $(COMPOSE_FILE) up -d
clean:
	@docker compose -f $(COMPOSE_FILE) down
	@rm -rf ./docs