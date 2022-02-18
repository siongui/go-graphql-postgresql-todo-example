package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/siongui/go-kit-gqlgen-postgres-todo-example/graph/model"
)

type TodoSearchRequest struct {
	T model.TodoSearchInput
	P model.PaginationInput
}

type CreateTodoRequest struct {
	T         model.CreateTodoInput
	CreatedBy string
}

type UpdateTodoRequest struct {
	Id        string
	T         model.UpdateTodoInput
	UpdatedBy string
}

type Resolver struct {
	Logger             log.Logger
	GetTodoEndpoint    endpoint.Endpoint
	TodoPagesEndpoint  endpoint.Endpoint
	TodoSearchEndpoint endpoint.Endpoint
	CreateTodoEndpoint endpoint.Endpoint
	UpdateTodoEndpoint endpoint.Endpoint
}
