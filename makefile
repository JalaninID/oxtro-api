include .env

gen-proto:
	rm -rf gen && buf generate
wire :
	wire gen app/injector

test :
	go test -v ./...
gow :
	@echo "Run server"
	export \
	DB_TYPE=$(DB_TYPE) \
	DB_USER=$(DB_USER) \
	DB_NAME=$(DB_NAME) \
	DB_PASS=$(DB_PASS) \
	DB_HOST=$(DB_HOST) \
	DB_PORT=$(DB_PORT) \
	JWT_TIME_DURATION=$(JWT_TIME_DURATION) \
	JWT_SECRET_KEY=$(JWT_SECRET_KEY) \
	&& gow run main.go

url=$(DB_TYPE)://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
migration-up:
	migrate -database "$(url)" -path $(DB_MIGRATION_PATH) up $(version)
	
migration-down:
	migrate -database "$(url)" -path $(DB_MIGRATION_PATH) down $(version)
	
migration-create:
	migrate create -ext sql -dir $(DB_MIGRATION_PATH) -seq $(name)

migration-force:
	migrate -database "$(url)" -path $(DB_MIGRATION_PATH) force $(version)

migration-version:
	migrate -database "$(url)" -path $(DB_MIGRATION_PATH) version