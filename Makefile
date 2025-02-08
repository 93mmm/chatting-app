# TODO: is $PROJECT_ROOT same as $PWD?

PROJECT_ROOT := $(shell pwd)

# Docker Compose
run_docker:
	@echo docker-compose up
	docker-compose --env-file .env -f docker/compose.yaml up --build

# Golang-Postgresql server
run_pgserver:
	@echo pgsql server up
	cd app/internal/services/pgserver && go mod tidy && cd cmd && CHATTING_APP_WORKING_DIRECTORY=$(PROJECT_ROOT) go run .

# Golang chatting-app server
run_chatting_app:
	@echo chatting-app server up
	cd app/internal/services/server && go mod tidy && cd cmd && CHATTING_APP_WORKING_DIRECTORY=$(PROJECT_ROOT) go run .

# Localhost server to open index.html and test api
run_localhost:
	@echo localhost up
	python3 -m http.server 8000
