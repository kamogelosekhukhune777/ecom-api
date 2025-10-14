BASE_IMAGE_NAME := ecomm
ECOM_APP        := ecom-api
VERSION         := 0.0.1
ECOM_IMAGE      := $(BASE_IMAGE_NAME)/$(ECOM_APP):$(VERSION)
# ===================================================================================================
# Building containers

build: ecom-api

ecom-api:
	docker build \
		-f zarf/docker/dockerfile.ecom \
		-t $(ECOM_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

# ===================================================================================================
# Docker Compose

compose-up:
	cd ./zarf/compose/ && docker compose -f docker_compose.yaml -p compose up -d

compose-build-up: build compose-up

compose-down:
	cd ./zarf/compose/ && docker compose -f docker_compose.yaml down

compose-logs:
	cd ./zarf/compose/ && docker compose -f docker_compose.yaml logs

# ===================================================================================================

run: 
	go run api/ecom/main.go

# ===================================================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor