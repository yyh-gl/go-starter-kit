export CGO_ENABLED=0

.DEFAULT_GOAL := help
.PHONY: help
help: ## Show help
	@echo '  see: https://git.dmm.com/dmm-app/pointclub-api'
	@echo ''
	@grep -E '^[%/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-22s\033[0m %s\n", $$1, $$2}'
	@echo ''

.PHONY: init
init: ## Initialize project
	cp $(shell pwd)/.scripts/pre-commit $(shell pwd)/.git/hooks/pre-commit
	chmod +x $(shell pwd)/.git/hooks/pre-commit

.PHONY: build
build: ## Build
	GOOS=linux GOARCH=amd64 go build -o ./cmd/api/bin/server ./cmd/api/main.go

.PHONY: lint
lint: ## Lint
	docker-compose exec  -T api golangci-lint run

.PHONY: test
test: ## Test
	docker-compose exec  -T api go test ./...
