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

PKGNAME=github.com/siongui/go-graphql-postgresql-todo-example
GQLGEN?=go run github.com/99designs/gqlgen
ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

run: fmt
	go run $(ALL_GO_SOURCES)

fmt:
	@go fmt *.go
	@go fmt config/*.go
	@go fmt graph/*.go
	@go fmt todo/*.go
	@go fmt tools/*.go

gqlinit:
	$(GQLGEN) init --verbose

schema_generate:
	$(GQLGEN) generate --verbose

local_dev_get_metrics:
	curl http://localhost:8080/metrics

modinit:
	go mod init $(PKGNAME)

modtidy:
	go mod tidy -go=1.16 && go mod tidy -go=1.17
