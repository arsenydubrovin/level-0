include .env
export

all: help

.PHONY: deps
# Update dependencies
deps: init
	go mod tidy -v

.PHONY: lint
# Lint the project
lint:
	golangci-lint run ./src/...

.PHONY: migrate-up
# Apply migrations
migrate-up:
	migrate -database "postgres://$(PG_USER):@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=disable" -path ./src/migrations up

.PHONY: migrate-down
# Undo migrations
migrate-down:
	migrate -database "postgres://$(PG_USER):@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=disable" -path ./src/migrations down

.PHONY: postgres-up
postgres-up:
	docker compose up -d postgres

.PHONY: nats-streaming-up
nats-streaming-up:
	docker compose up -d nats-streaming

.PHONY: run
# Run the application and the database container
run: postgres-up nats-streaming-up migrate-up
	air

.PHONY: init
# Initialize the repository for development
init: install-gofumpt install-air install-golangci-lint install-precommit install-migrate
ifeq (,$(wildcard .git/hooks/pre-commit))
	pre-commit install
endif

.PHONY: install-gofumpt
install-gofumpt:
ifeq (, $(shell which gofumpt))
	echo "Installing gofumpt..."
	go install mvdan.cc/gofumpt@latest
endif

.PHONY: install-air
install-air:
ifeq (, $(shell which air))
	echo "Installing air..."
	go install github.com/cosmtrek/air@latest
endif

.PHONY: install-precommit
install-precommit:
ifeq (, $(shell which pre-commit))
	echo "Installing pre-commit..."
	python3 -m pip install pre-commit
endif

.PHONY: install-golangci-lint
install-golangci-lint:
ifeq (, $(shell which golangci-lint))
	echo "Installing golangci-lint..."
	$(shell curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.55.2)
endif

.PHONY: install-migrate
install-migrate:
ifeq (, $(shell which migrate))
	echo "Installing migrate..."
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
endif

.PHONY: help
# Show this help
help:
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t
