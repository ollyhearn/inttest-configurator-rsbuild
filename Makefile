BASE_DIR = $(shell pwd)

start.db:
	cd ./deployments && docker compose up -d db

start.local:
	DATABASE_URL="postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable&" AUTO_MIGRATE=true \
	go run ./cmd/configurator

prepare-swag-image:
	docker build -f deployments/Dockerfile-swag . -t swag-image:1.8.4

swag.gen:
	docker run --rm  -it  -v "$(BASE_DIR):/dockerdev" swag-image:1.8.4  /bin/bash -c \
		"swag init --overridesFile .swaggo --generatedTime --parseDependency --parseInternal --parseDepth 1 -g cmd/configurator/main.go"

## swag.generate.ui.api: собрать api на основе swagger для UI
swag.generate.ui.api:
	docker run --rm -v "$(BASE_DIR):/local" openapitools/openapi-generator-cli:v6.0.1 generate \
		-i /local/docs/swagger.yaml \
		-g javascript \
		-o /local/frontend/src/api
