DOCS_DIR=api/docs
DOCS_FILE=sso.json

DB_URL=postgresql://postgres:postgres@localhost:5555/postgres?sslmode=disable



generate-docs:
	mkdir -p $(DOCS_DIR) && \
	protoc -I $(PROTO_DIR) \
		-I internal/pkg/googleapis \
		--openapiv2_out=$(DOCS_DIR) \
		--openapiv2_opt=logtostderr=true,allow_merge=true \
		$(PROTO_DIR)/*.proto


create-db-image:
	docker run --name cifra-sso -p 5555:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

migrate-up:
	migrate -path internal/db/migration -database $(DB_URL) -verbose up

migrate-down:
	migrate -path internal/db/migration -database $(DB_URL) -verbose down

generate-sqlc:
	sqlc generate

build-server:
	go build -o main cmd/sso/main.go

run-server: build-server
	go run cmd/sso/main.go

run-server-dev:
	go build -o main cmd/sso/main.go && go run cmd/sso/main.go