.PHONY: default run run-woth-docs build run-executable test docs stack clean stop

#Variables
APP_NAME=ps-tag-onboarding-go
COMPOSE_FILE=build/.docker/docker-compose.yaml
MAIN_FILE=cmd/ps-tag-onboarding/main.go

# Tasks
default: run-with-docs

run:
	@docker compose -f $(COMPOSE_FILE) up -d
run-with-docs:
	@swag init -g cmd/ps-tag-onboarding/main.go --parseInternal --output ./api
	@docker compose -f $(COMPOSE_FILE) up -d
build:
	@go build -o $(APP_NAME) $(MAIN_FILE)
run-executable: build
	@./$(APP_NAME)
test:
	@go test ./...
docs:
	@swag init -g cmd/ps-tag-onboarding/main.go --parseInternal --output ./api
clean:
	@docker compose -f $(COMPOSE_FILE) down
	@rm -rf ./api
stop:
	@docker compose -f $(COMPOSE_FILE) down