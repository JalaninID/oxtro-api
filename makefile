include .env

gen-proto:
	rm -rf gen && buf generate oxtro-proto

gen-proto-plugin:
	rm -rf plugins/sample_crm/gen && buf generate plugins/sample_crm/proto --template plugins/sample_crm/buf.gen.yaml

gen-proto-all: gen-proto gen-proto-plugin
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