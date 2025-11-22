GO_PACKAGES ?= $(shell go list ./... | grep -v 'examples\|qtest\|mock\|cmd\|test')

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

integration-test:
	go test ./module/configuration/test/integration -v

coverage:
	go test -race -cover -coverprofile=coverage.out -json ${GO_PACKAGES} > ./UT-report_coverage.json
	go tool cover -func=coverage.out

cover:
	go tool cover -html=coverage.out

build-image:
	docker build -f build/rest/Dockerfile -t configs-service .

run-image:
	docker run -d --name configs-service-app -p 8080:8080 configs-service