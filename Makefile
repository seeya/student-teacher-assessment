include .env

$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))

migrate.up:
	docker run -v "${PWD}/platform/migrations":/migrations --network host migrate/migrate \
		-path=/migrations/ \
		-database "mysql://root:${MYSQL_ROOT_PASSWORD}@tcp(127.0.0.1:3306)/${MYSQL_DATABASE}" \
		up

migrate.down:
	docker run -v "${PWD}/platform/migrations":/migrations --network host migrate/migrate \
		-path=/migrations/ \
		-database "mysql://root:${MYSQL_ROOT_PASSWORD}@tcp(127.0.0.1:3306)/${MYSQL_DATABASE}" \
		down -all	