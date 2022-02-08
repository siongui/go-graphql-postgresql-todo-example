# Check if a program exists from a Makefile
# https://stackoverflow.com/a/34756868
CMDGO := $(shell command -v go 2> /dev/null)
ifndef GOROOT
ifndef CMDGO
export GOROOT=$(realpath ../go/)
export PATH := $(GOROOT)/bin:$(PATH)
endif
endif
ifndef GOPATH
export GOPATH=$(realpath ./tools/)
export PATH := $(GOPATH)/bin:$(PATH)
endif

PKGNAME=github.com/siongui/go-kit-gqlgen-postgres-todo-example
MIGRATIONS_DIR=$(CURDIR)/migrations/
GQLGEN?=go run github.com/99designs/gqlgen
ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

run: fmt
	go run $(ALL_GO_SOURCES)

fmt:
	@go fmt *.go
	@go fmt config/*.go
	@go fmt graph/*.go
	@go fmt graph/scalar/*.go
	@go fmt todo/*.go
	@go fmt todo/tododb/*.go
	@go fmt tools/*.go
	@go fmt tools/migrate/*.go

gqlinit:
	$(GQLGEN) init --verbose

schema_generate:
	$(GQLGEN) generate --verbose

golangci_lint:
	golangci-lint run main.go
	golangci-lint run ./config/...
	golangci-lint run ./todo/...
	golangci-lint run ./graph/...
	golangci-lint run ./tools/migrate/...

graphql_schema_lint:
	npx graphql-schema-linter

database_migrations: fmt
	go run tools/migrate/main.go -ssl=false -dir=$(MIGRATIONS_DIR)

local_dev_get_metrics:
	curl http://localhost:3005/metrics

modinit:
	go mod init $(PKGNAME)

modtidy:
	go mod tidy -go=1.16 && go mod tidy -go=1.17

install_golangci_lint:
	@echo "\033[92mInstall golangci-lint ...\033[0m"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.0
	golangci-lint --version
