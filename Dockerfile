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
COPY ./config/ ./config/
COPY ./todo/ ./todo/
COPY ./tools/ ./tools/
COPY gqlgen.yml ./
COPY ./graph/resolver.go ./graph/resolver.go
COPY ./graph/schema.graphqls ./graph/schema.graphqls
COPY ./graph/schema.resolvers.go ./graph/schema.resolvers.go
COPY ./graph/scalar/ ./graph/scalar/
COPY Makefile ./Makefile

RUN apk add build-base
RUN make schema_generate
RUN go build -o /todo_server

EXPOSE 3005

CMD [ "/todo_server" ]
