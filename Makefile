# File: Makefile

AUTH_SERVICE_NAME=lifekost/auth-service
AUTH_SERVICE_PATH=services/auth-services

.PHONY: build-auth run-auth stop-auth tidy-auth

build-auth:
	docker build -f $(AUTH_SERVICE_PATH)/Dockerfile -t $(AUTH_SERVICE_NAME) .

run-auth:
	docker run -d --name auth-service -p 8081:8081 $(AUTH_SERVICE_NAME)

stop-auth:
	docker stop auth-service || true
	docker rm auth-service || true

tidy-auth:
	cd $(AUTH_SERVICE_PATH) && go mod tidy