package todo

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/siongui/go-graphql-postgresql-todo-example/graph"
	"github.com/siongui/go-graphql-postgresql-todo-example/graph/model"
)

func makeGetTodoEndpoint(svc TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := request.(string)
		return svc.GetTodo(id)
	}
}

func makeTodoPagesEndpoint(svc TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.TodoPages(request.(model.PaginationInput))
	}
}

func makeTodoSearchEndpoint(svc TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(graph.TodoSearchRequest)
		return svc.TodoSearch(req.T, req.P)
	}
}

func makeCreateTodoEndpoint(svc TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.CreateTodo(request.(model.TodoInput))
	}
}

func makeUpdateTodoEndpoint(svc TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(graph.UpdateTodoRequest)
		return svc.UpdateTodo(req.Id, req.T)
	}
}
