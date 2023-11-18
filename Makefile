DATABASE_URL="mysql://root:root@tcp\(localhost:3306\)/test"

run:
	@go run cmd/app/main.go

create-migration:
	@migrate create -ext sql -dir ./migrations $(MIGRATION_NAME)

migrate:
	@migrate -path migrations -database "$(DATABASE_URL)" up

destroy-db:
	@migrate -path migrations -database "$(DATABASE_URL)" down