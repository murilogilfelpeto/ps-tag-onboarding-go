.PHONY: default run build test docs stack clean

#Variables
APP_NAME=ps-tag-onboarding-go
COMPOSE_FILE=build/.docker/stack.yaml

# Tasks
default: run-with-docs

run:
	@docker compose -f $(COMPOSE_FILE) up -d
run-with-docs:
	@swag init -g cmd/ps-tag-onboarding/main.go --parseInternal --output ./api
	@docker compose -f $(COMPOSE_FILE) up -d
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init -g cmd/ps-tag-onboarding/main.go --parseInternal --output ./api
stack:
	@docker compose -f $(COMPOSE_FILE) up -d
clean:
	@docker compose -f $(COMPOSE_FILE) down
	@rm -rf ./api