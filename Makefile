export CGO_ENABLED=0

.PHONY: lint
lint:
	docker-compose exec  -T api golangci-lint run

.PHONY: test
test:
	docker-compose exec  -T api go test ./...

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o ./cmd/api/bin/server ./cmd/api/main.go
