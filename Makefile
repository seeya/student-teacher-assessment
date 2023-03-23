include .env

$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))


app.run:
	go run main.go

migrate.up:
	docker run -v "${PWD}/platform/migrations":/migrations --network host migrate/migrate \
		-path=/migrations/ \
		-database "mysql://root:${MYSQL_ROOT_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" \
		up

migrate.down:
	docker run -v "${PWD}/platform/migrations":/migrations --network host migrate/migrate \
		-path=/migrations/ \
		-database "mysql://root:${MYSQL_ROOT_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" \
		down -all	