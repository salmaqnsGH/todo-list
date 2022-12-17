migrateup:
	migrate -database "mysql://user:password@tcp(127.0.0.1:3306)/todolist" -path db/migrations up

migratedown:
	migrate -database "mysql://user:password@tcp(127.0.0.1:3306)/todolist" -path db/migrations down

.PHONY: migrateup migratedown