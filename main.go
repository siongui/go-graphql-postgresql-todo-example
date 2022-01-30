package main

import (
	"context"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/siongui/go-graphql-postgresql-todo-example/graph"
	"github.com/siongui/go-graphql-postgresql-todo-example/graph/generated"
)

const defaultPort = "8080"

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

func transportLoggingMiddleware(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logger := log.NewLogfmtLogger(os.Stderr)

	svc := todoService{}

	var getTodoEndpoint endpoint.Endpoint
	getTodoEndpoint = makeGetTodoEndpoint(svc)
	getTodoEndpoint = transportLoggingMiddleware(log.With(logger, "method", "getTodo"))(getTodoEndpoint)

	var todoPagesEndpoint endpoint.Endpoint
	todoPagesEndpoint = makeTodoPagesEndpoint(svc)
	todoPagesEndpoint = transportLoggingMiddleware(log.With(logger, "method", "TodoPages"))(todoPagesEndpoint)

	var todoSearchEndpoint endpoint.Endpoint
	todoSearchEndpoint = makeTodoSearchEndpoint(svc)
	todoSearchEndpoint = transportLoggingMiddleware(log.With(logger, "method", "TodoSearch"))(todoSearchEndpoint)

	var createTodoEndpoint endpoint.Endpoint
	createTodoEndpoint = makeCreateTodoEndpoint(svc)
	createTodoEndpoint = transportLoggingMiddleware(log.With(logger, "method", "createTodo"))(createTodoEndpoint)

	var updateTodoEndpoint endpoint.Endpoint
	updateTodoEndpoint = makeUpdateTodoEndpoint(svc)
	updateTodoEndpoint = transportLoggingMiddleware(log.With(logger, "method", "updateTodo"))(updateTodoEndpoint)

	graphQLHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			GetTodoEndpoint:    getTodoEndpoint,
			TodoPagesEndpoint:  todoPagesEndpoint,
			TodoSearchEndpoint: todoSearchEndpoint,
			CreateTodoEndpoint: createTodoEndpoint,
			UpdateTodoEndpoint: updateTodoEndpoint,
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graphQLHandler)

	logger.Log("msg", "connect to http://localhost:"+port+"/ for GraphQL playground")
	logger.Log("err", http.ListenAndServe(":"+port, nil))
}
