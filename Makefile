GO_PACKAGES ?= $(shell go list ./... | grep -v 'examples\|qtest\|mock')

create-migrations:
	migrate create -ext sql -dir ./module/configuration/db/migration -seq $(name)

migrate:
	go run ./cmd/migrate/main.go up

migrate-down:
	go run ./cmd/migrate/main.go down

inject-users:
	go run ./cmd/inject/main.go

run:
	go run ./cmd/rest/main.go

test:
	go test -race -v ${GO_PACKAGES}

coverage:
	go test -race -cover -coverprofile=coverage.out -json ${GO_PACKAGES} > ./UT-report_coverage.json
	go tool cover -func=coverage.out