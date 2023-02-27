migrate:
	migrate -database postgres://rootuser:nosecret@localhost:5432/billodb?sslmode=disable -path database/postgres/migration up

downmigrate:
	migrate -database postgres://rootuser:nosecret@localhost:5432/billodb?sslmode=disable -path database/postgres/migration down