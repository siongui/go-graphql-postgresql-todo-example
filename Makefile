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
	@go fmt tools/gqlclient/*.go
	@go fmt tools/migrate/filesystem/*.go
	@go fmt tools/migrate/iofs/*.go

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
	golangci-lint run ./tools/gqlclient/...

graphql_schema_lint:
	npx graphql-schema-linter

MIGRATE_TOOL_DIR=$(CURDIR)/tools/migrate/iofs
TMP_MIGRATIONS=$(MIGRATE_TOOL_DIR)/migrations
database_migrations_iofs: fmt
	@echo "\033[92mMigration Source: iofs ...\033[0m"
	[ -d $(TMP_MIGRATIONS) ] || rm -rf $(TMP_MIGRATIONS)
	cp -r $(MIGRATIONS_DIR) $(TMP_MIGRATIONS)
	cd $(MIGRATE_TOOL_DIR); go build -o $(GOPATH)/bin/migrate
	migrate -ssl=false

database_migrations: fmt
	@echo "\033[92mMigration Source: filesystem ...\033[0m"
	go run tools/migrate/filesystem/main.go -ssl=false -dir=$(MIGRATIONS_DIR)

run_gqlclient:
	go run tools/gqlclient/main.go

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
