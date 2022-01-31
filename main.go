package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-kit/kit/log"
	"github.com/siongui/go-graphql-postgresql-todo-example/todo"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logger := log.NewLogfmtLogger(os.Stderr)

	var svc todo.TodoService
	svc = todo.NewService()
	svc = todo.NewLoggingMiddleware(logger, svc)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", todo.MakeGraphQLHandler(svc, logger))

	logger.Log("msg", "connect to http://localhost:"+port+"/ for GraphQL playground")
	logger.Log("err", http.ListenAndServe(":"+port, nil))
}
