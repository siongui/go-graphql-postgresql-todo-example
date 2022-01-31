# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.17-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY ./todo/ ./todo/
COPY ./tools/ ./tools/
COPY gqlgen.yml ./
COPY ./graph/resolver.go ./graph/resolver.go
COPY ./graph/schema.graphqls ./graph/schema.graphqls
COPY ./graph/schema.resolvers.go ./graph/schema.resolvers.go

RUN apk add build-base
RUN go run github.com/99designs/gqlgen generate --verbose
RUN go build -o /todo_server

EXPOSE 8080

CMD [ "/todo_server" ]
