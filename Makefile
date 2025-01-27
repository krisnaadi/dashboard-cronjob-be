include .env

build:
	@go build -v -o cronjob-be ./cmd/*.go

run:
	@echo "RUN cronjob-be..."
	make build
	@./cronjob-be

migrate:
	@migrate -database ${DB_MIGRATION_CONNECTION} -path database/migrations up

migrate-rollback:
	@migrate -database ${DB_MIGRATION_CONNECTION} -path database/migrations down

seed:
	@migrate -database ${DB_SEEDER_CONNECTION} -path database/seeders up

seed-rollback:
	@migrate -database ${DB_SEEDER_CONNECTION} -path database/seeders down
