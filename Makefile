run:
	go run main.go

migrateup:
	migrate -path db/migrations -database "mysql://user:password@tcp(127.0.0.1:3306)/todolist" -verbose up

migratedown:
	migrate -path db/migrations -database "mysql://user:password@tcp(127.0.0.1:3306)/todolist" -verbose down

docker-build:
	docker build --tag todo-list .

docker-compose-up:
	docker compose up -d

docker-run-local:
	docker run todo-list

docker-run:
	docker run -e MYSQL_HOST=172.17.0.1 -e MYSQL_USER=user -e MYSQL_PASSWORD=password -e MYSQL_DBNAME=todolist -p 8090:3030 todo-list

.PHONY: migrateup migratedown