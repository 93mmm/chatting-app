# ========================================================================
# |                        Set working directory                         |
# ========================================================================
PROJECT_ROOT := $(shell pwd)



# ========================================================================
# |                         Run Docker commands                          |
# ========================================================================
run_app:
	docker-compose --env-file .env -f docker/compose.yaml up --build db db-service server

run_app_without_rebuild:
	docker-compose --env-file .env -f docker/compose.yaml up



# ========================================================================
# |                  Run tests inside Docker container                   |
# ========================================================================
run_tests:
	docker-compose --env-file .env -f docker/compose.yaml up --build tester

test:
	cd app/internal/tests/runners && CHATTING_APP_WORKING_DIRECTORY=$(PROJECT_ROOT) go test ./... -v



# ========================================================================
# |                       Launch app inside Docker                       |
# ========================================================================
pgserver:
	cd app/internal/services/pgserver && go mod tidy && cd cmd && CHATTING_APP_WORKING_DIRECTORY=$(PROJECT_ROOT) go run .

chatting_app:
	cd app/internal/services/goserver && go mod tidy && cd cmd && CHATTING_APP_WORKING_DIRECTORY=$(PROJECT_ROOT) go run .



# ========================================================================
# |                         Run pseudo-frontend                          |
# ========================================================================
# Localhost server to open index.html and test api
localhost:
	python3 -m http.server 8000
