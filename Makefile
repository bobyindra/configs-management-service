create-migrations:
	migrate create -ext sql -dir ./db/migrations -seq $(name)

migrate:
	go run ./cmd/migrate/main.go up

migrate-down:
	go run ./cmd/migrate/main.go down

inject-users:
	go run ./cmd/inject/main.go

run:
	go run ./cmd/rest/main.go