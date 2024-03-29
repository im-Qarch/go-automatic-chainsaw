postgres:
	docker run --name bank_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:latest

createdb:
	docker exec -it bank_postgres createdb --username=root --owner=root bank_db
 
dropdb:
	# This command will terminate all connections to the "bank_db" database in your PostgreSQL Docker container.
	docker exec -it bank_postgres psql -U root -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = 'bank_db' AND pid <> pg_backend_pid();"
	# This should allow you to successfully drop the database without encountering the error about active connections.
	docker exec -it bank_postgres dropdb bank_db

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc