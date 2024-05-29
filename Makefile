.PHONY: up down create

DATABASE_URL ?= POSTGRES_URL
MIGRATIONS_PATH ?= ./internal/database/migrations

PRIVATE_BUILD=go build -o ./bin/eccommerce/main ./cmd/eccommerce/main.go
PRIVATE_COMMAND=./bin/eccommerce/main

PUBLIC_BUILD=go build -o ./bin/public/main ./cmd/public/main.go

# go get -u -d github.com/golang-migrate/migrate/cmd/migrate
# linux system only
setup-nginx:
	sudo chmod +x ./pkg/scripts/nginx.sh
	sudo ./pkg/scripts/nginx.sh

setup-private:
	${PRIVATE_BUILD}
	sudo chmod +x ./pkg/scripts/private.sh
	sudo ./pkg/scripts/private.sh

setup-public:
	${PUBLIC_BUILD}
	sudo chmod +x ./pkg/scripts/public.sh
	sudo ./pkg/scripts/public.sh

restart:
	sudo chmod +x ./pkg/scripts/rebuild.sh
	sudo ./pkg/scripts/rebuild.sh

up:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_PATH) up $(id)

down:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_PATH) down

create:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) $(name)

force:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_PATH) force $(id)

version:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_PATH) version


start-private:
	${PRIVATE_BUILD}
	${PRIVATE_COMMAND}
	

start-public:
	go run ./cmd/public/main.go

build:
	go build -o ./bin/eccommerce/main ./cmd/eccommerce/main.go 
	go build -o ./bin/public/main ./cmd/public/main.go 
