include .env

$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))


app.run:
	go run main.go

seed:
	curl http://localhost:3001/api/seed

migrate.up:
	docker run -v "${PWD}/platform/migrations":/migrations --network host migrate/migrate \
		-path=/migrations/ \
		-database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(127.0.0.1:${MYSQL_PORT})/${MYSQL_DATABASE}" \
		up

migrate.down:
	docker run -v "${PWD}/platform/migrations":/migrations --network host migrate/migrate \
		-path=/migrations/ \
		-database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(127.0.0.1:${MYSQL_PORT})/${MYSQL_DATABASE}" \
		down -all	

build:
	docker-compose up --build

up:
	docker-compose up

down:
	docker-compose down

prune.volume: 
	docker volume prune

test:
	go clean -testcache && go test -p 1 -v -cover ./...

.PHONY: migrate.up migrate.down test 