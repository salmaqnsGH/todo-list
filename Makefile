run:
	go run main.go

migrateup:
	migrate -path db/migrations -database "mysql://user:password@tcp(127.0.0.1:3306)/todolist" -verbose up

migratedown:
	migrate -path db/migrations -database "mysql://user:password@tcp(127.0.0.1:3306)/todolist" -verbose down

docker-build:
	docker build --tag todo-list .

docker-run-local:
	docker run todo-list

.PHONY: migrateup migratedown