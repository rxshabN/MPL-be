run:
	go run cmd/main.go

postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_PASSWORD=Orchid7890 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=postgres mpl-db

dropdb:
	docker exec -it postgres12 dropdb 
	