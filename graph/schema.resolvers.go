package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/siongui/go-graphql-postgresql-todo-example/graph/generated"
	"github.com/siongui/go-graphql-postgresql-todo-example/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.TodoInput) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id int, input model.TodoInput) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTodo(ctx context.Context, id int) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TodoPages(ctx context.Context, paginationInput model.PaginationInput) (*model.TodoPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TodoSearch(ctx context.Context, input model.TodoSearchInput, paginationInput model.PaginationInput) (*model.TodoPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
