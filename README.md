1)docker pull postgres

2)docker run -d \
--name my-postgres \
-p 5433:5432 \
-e POSTGRES_USER=myuser \
-e POSTGRES_PASSWORD=mypassword \
-e POSTGRES_DB=mydatabase \
postgres:latest
