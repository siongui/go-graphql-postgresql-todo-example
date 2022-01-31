package todo

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/siongui/go-graphql-postgresql-todo-example/graph"
	"github.com/siongui/go-graphql-postgresql-todo-example/graph/generated"
)

func transportLoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}

func MakeGraphQLHandler(svc TodoService, logger log.Logger) http.Handler {
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

	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			GetTodoEndpoint:    getTodoEndpoint,
			TodoPagesEndpoint:  todoPagesEndpoint,
			TodoSearchEndpoint: todoSearchEndpoint,
			CreateTodoEndpoint: createTodoEndpoint,
			UpdateTodoEndpoint: updateTodoEndpoint,
		},
	}))
}
