start.db:
	cd ./deployments && docker compose up -d db

start.local:
	DATABASE_URL="postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable&" AUTO_MIGRATE=true \
	go run ./cmd/configurator