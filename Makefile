DB_URL=postgresql://postgres:postgres@localhost:5555/postgres?sslmode=disable
SWAGGER_CODEGEN := java -jar /home/trpdjke/go/src/github.com/cifra-city/rest-sso/swagger-codegen-cli.jar

OPENAPI_GENERATOR := java -jar ./openapi-generator-cli.jar

generate-models:
	$(OPENAPI_GENERATOR) generate -i docs/api.yaml -g go -o ./docs/web --additional-properties=packageName=resources
	mkdir -p resources
	find docs/web -name '*.go' -exec mv {} resources/ \;


create-db-image:
	docker run --name cifra-sso -p 5555:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

migrate-up:
	migrate -path internal/db/migration -database $(DB_URL) -verbose up

migrate-down:
	migrate -path internal/db/migration -database $(DB_URL) -verbose down

generate-sqlc:
	sqlc generate

build-server:
	go build -o main main.go

run-server: build-server
	go run main.go
