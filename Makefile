DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgres-12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres-12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	docker run -v $(shell pwd):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/nikit34/template_backend/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/nikit34/template_backend/worker TaskDistributor

protoc:
	rm -f pb/*.go
	rm -rf doc/swagger/*.swagger.json
	protoc \
	--proto_path=proto \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=. \
	--openapiv2_out=doc/swagger \
	--openapiv2_opt=allow_merge=true,merge_file_name=template_backend \
	proto/*.proto
	statik -f -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY:
	postgres
	createdb
	dropdb
	migrateup
	migratedown
	migrateup1
	migratedown1
	new_migration
	db_docs
	db_schema
	sqlc
	test
	server
	mock
	protoc
	evans
	redis
