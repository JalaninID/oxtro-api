include .env

gen-proto:
	rm -rf gen && buf generate
wire :
	wire gen app/injector

test :
	go test -v ./...
run :
	sh -c 'set -a; . ./.env; set +a; gow run main.go'

url=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)

migration-up:
	migrate -database "$(url)" -path ./migrations/ up $(version)
	
migration-down:
	migrate -database "$(url)" -path ./migrations/ down $(version)
	
migration-create:
	migrate create -ext sql -dir ./migrations/ -seq $(name)

migration-force:
	migrate -database "$(url)" -path ./migrations/ force $(version)

migration-version:
	migrate -database "$(url)" -path ./migrations/ version