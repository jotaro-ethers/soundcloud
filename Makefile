postgres:
	docker run --name my-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -d postgres
createdb: 
	docker exec -it my-postgres createdb --username=root --owner=root soundcloud
dropdb:
	docker exec -it my-postgres dropdb soundcloud
migrateup:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@172.17.0.2:5432/soundcloud?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@172.17.0.2:5432/soundcloud?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...	

		
