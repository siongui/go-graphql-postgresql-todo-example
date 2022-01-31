package todo

import (
	"github.com/siongui/go-graphql-postgresql-todo-example/graph/model"
)

type TodoService interface {
	GetTodo(int) (*model.Todo, error)
	TodoPages(model.PaginationInput) (*model.TodoPagination, error)
	TodoSearch(model.TodoSearchInput, model.PaginationInput) (*model.TodoPagination, error)
	CreateTodo(model.TodoInput) (*model.Todo, error)
	UpdateTodo(int, model.TodoInput) (*model.Todo, error)
}

type todoService struct{}

func (s *todoService) GetTodo(id int) (t *model.Todo, err error) {
	t = &model.Todo{}
	return
}

func (s *todoService) TodoPages(pi model.PaginationInput) (tp *model.TodoPagination, err error) {
	tp = &model.TodoPagination{}
	return
}

func (s *todoService) TodoSearch(tsi model.TodoSearchInput, pi model.PaginationInput) (tp *model.TodoPagination, err error) {
	tp = &model.TodoPagination{}
	return
}

func (s *todoService) CreateTodo(ti model.TodoInput) (t *model.Todo, err error) {
	t = &model.Todo{}
	return
}

func (s *todoService) UpdateTodo(id int, ti model.TodoInput) (t *model.Todo, err error) {
	t = &model.Todo{}
	return
}

func NewService() TodoService {
	return &todoService{}
}

// ServiceMiddleware is a chainable behavior modifier for TodoService.
type ServiceMiddleware func(TodoService) TodoService
