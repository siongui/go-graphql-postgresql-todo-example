package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetTodoEndpoint(svc TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := request.(int)
		return svc.GetTodo(id)
	}
}

func makeTodoPagesEndpoint(svc TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.TodoPages(request.(PaginationInput))
	}
}

func makeTodoSearchEndpoint(svc TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(todoSearchRequest)
		return svc.TodoSearch(req.T, req.P)
	}
}

func makeCreateTodoEndpoint(svc TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.CreateTodo(request.(TodoInput))
	}
}

func makeUpdateTodoEndpoint(svc TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateTodoRequest)
		return svc.UpdateTodo(req.Id, req.T)
	}
}

type todoSearchRequest struct {
	T TodoSearchInput
	P PaginationInput
}

type updateTodoRequest struct {
	Id int
	T  TodoInput
}
