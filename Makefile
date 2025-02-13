PROJECT_ROOT := $(shell pwd)

# Docker Compose
docker:
	@echo docker-compose up
	docker-compose --env-file .env -f docker/compose.yaml up --build

docker_without_rebuild:
	@echo docker-compose up
	docker-compose --env-file .env -f docker/compose.yaml up

# Golang-Postgresql server
pgserver:
	@echo pgsql server up
	cd app/internal/services/pgserver && go mod tidy && cd cmd && CHATTING_APP_WORKING_DIRECTORY=$(PROJECT_ROOT) go run .

# Golang chatting-app server
chatting_app:
	@echo chatting-app server up
	cd app/internal/services/goserver && go mod tidy && cd cmd && CHATTING_APP_WORKING_DIRECTORY=$(PROJECT_ROOT) go run .

# Localhost server to open index.html and test api
localhost:
	@echo localhost up
	python3 -m http.server 8000

# Run tests
test:
	@echo testing pgserver
	cd app/internal/services/pgserver && go test ./... -v
	@echo testing goserver
	cd app/internal/services/goserver && go test ./... -v
