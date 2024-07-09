all:
	docker-compose up --build
build:
	go mod tidy
	go run cmd/main.go
postgresrun:
	docker run -d \
    --name my-postgres \
    -p 5432:5432 \
    -e POSTGRES_USER=myuser \
    -e POSTGRES_PASSWORD=mypassword \
    -e POSTGRES_DB=mydatabase \
    postgres:latest