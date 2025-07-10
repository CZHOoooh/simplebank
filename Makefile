postgres:
	docker run --name postgres-container --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e TZ=Asia/Shanghai -d postgres:latest

createdb:
	docker exec -it postgres-container createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-container dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root_czh:Curry86835467@pgm-wz9m87577cw5l3whko.pg.rds.aliyuncs.com:5432/simple_bank" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	 mockgen -package mockdb --destination db/mock/store.go simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock