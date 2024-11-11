DB_URL=postgresql://postgres:postgres@localhost:5555/postgres?sslmode=disable
SWAGGER_CODEGEN := java -jar /home/trpdjke/go/src/github.com/cifra-city/rest-sso/swagger-codegen-cli.jar

SWAGGER_CODEGEN := java -jar /home/trpdjke/go/src/github.com/cifra-city/rest-sso/swagger-codegen-cli.jar

generate-models:
	$(SWAGGER_CODEGEN) generate -i docs/api.yaml -l go -o ./docs/web
	mkdir -p models
	find docs/web -name '*.go' -exec mv {} models/ \;

create-db-image:
	docker run --name cifra-sso -p 5555:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

migrate-up:
	migrate -path internal/db/migration -database $(DB_URL) -verbose up

migrate-down:
	migrate -path internal/db/migration -database $(DB_URL) -verbose down

generate-sqlc:
	sqlc generate

