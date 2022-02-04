package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/siongui/go-graphql-postgresql-todo-example/graph/model"
)

type TodoSearchRequest struct {
	T model.TodoSearchInput
	P model.PaginationInput
}

type UpdateTodoRequest struct {
	Id string
	T  model.UpdateTodoInput
}

type Resolver struct {
	GetTodoEndpoint    endpoint.Endpoint
	TodoPagesEndpoint  endpoint.Endpoint
	TodoSearchEndpoint endpoint.Endpoint
	CreateTodoEndpoint endpoint.Endpoint
	UpdateTodoEndpoint endpoint.Endpoint
}
