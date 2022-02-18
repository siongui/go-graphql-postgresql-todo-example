package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/siongui/go-kit-gqlgen-postgres-todo-example/graph/generated"
	"github.com/siongui/go-kit-gqlgen-postgres-todo-example/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	// TODO: createdby is email address of login user. use fixed value now.
	createdby := "my@example.com"
	t, err := r.CreateTodoEndpoint(ctx, CreateTodoRequest{T: input, CreatedBy: createdby})
	return t.(*model.Todo), err
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, input model.UpdateTodoInput) (*model.Todo, error) {
	// TODO: updatedby is email address of login user. use fixed value now.
	updatedby := "updator@example.com"
	t, err := r.UpdateTodoEndpoint(ctx, UpdateTodoRequest{Id: id, T: input, UpdatedBy: updatedby})
	return t.(*model.Todo), err
}

func (r *queryResolver) GetTodo(ctx context.Context, id string) (*model.Todo, error) {
	t, err := r.GetTodoEndpoint(ctx, id)
	return t.(*model.Todo), err
}

func (r *queryResolver) TodoPages(ctx context.Context, paginationInput model.PaginationInput) (*model.TodoPagination, error) {
	t, err := r.TodoPagesEndpoint(ctx, paginationInput)
	return t.(*model.TodoPagination), err
}

func (r *queryResolver) TodoSearch(ctx context.Context, input model.TodoSearchInput, paginationInput model.PaginationInput) (*model.TodoPagination, error) {
	t, err := r.TodoSearchEndpoint(ctx, TodoSearchRequest{T: input, P: paginationInput})
	return t.(*model.TodoPagination), err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
