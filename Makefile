ifndef GOROOT
export GOROOT=$(realpath ../go/)
export PATH := $(GOROOT)/bin:$(PATH)
endif
ifndef GOPATH
export GOPATH=$(realpath ./graph/)
export PATH := $(GOPATH)/bin:$(PATH)
endif

PKGNAME=github.com/siongui/go-graphql-postgresql-todo-example
GQLGEN?=go run github.com/99designs/gqlgen
ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

run: fmt
	go run $(ALL_GO_SOURCES)

fmt:
	@go fmt *.go
	@go fmt tools/*.go
	@go fmt graph/*.go

gqlinit:
	$(GQLGEN) init --verbose

schema_generate:
	$(GQLGEN) generate --verbose

modinit:
	go mod init $(PKGNAME)

modtidy:
	go mod tidy -go=1.16 && go mod tidy -go=1.17
