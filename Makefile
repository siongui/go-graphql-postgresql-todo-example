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

run:
	go run server.go

fmt:
	go fmt server.go
	go fmt tools.go

gqlinit:
	$(GQLGEN) init --verbose

schema_generate:
	$(GQLGEN) generate --verbose

modinit:
	go mod init $(PKGNAME)

modtidy:
	go mod tidy
