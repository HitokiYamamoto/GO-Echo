.PHONY: fmt

fmt:
	go fmt ./...

lint:
	golangci-lint run

migration:
	GO_ENV=dev go run migrate/migrate.go