package todo

import (
	"errors"
	"strconv"
	"time"

	"github.com/siongui/go-graphql-postgresql-todo-example/graph/model"
	"github.com/siongui/go-graphql-postgresql-todo-example/todo/tododb"
)

type TodoService interface {
	GetTodo(string) (*model.Todo, error)
	TodoPages(model.PaginationInput) (*model.TodoPagination, error)
	TodoSearch(model.TodoSearchInput, model.PaginationInput) (*model.TodoPagination, error)
	CreateTodo(model.TodoInput) (*model.Todo, error)
	UpdateTodo(string, model.TodoInput) (*model.Todo, error)
}

type todoService struct {
	store tododb.TodoStore
}

func (s *todoService) GetTodo(id string) (t *model.Todo, err error) {
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
	td := tododb.Todo{}
	if ti.ContentCode != nil {
		td.ContentCode = *ti.ContentCode
	}
	if ti.ContentName != nil {
		td.ContentName = *ti.ContentName
	}
	if ti.Description != nil {
		td.Description = *ti.Description
	}
	if ti.StartDate != nil {
		startDate, err := time.Parse(time.RFC3339, *ti.StartDate)
		if err != nil {
			return t, err
		}
		td.StartDate = startDate
	}
	if ti.EndDate != nil {
		endDate, err := time.Parse(time.RFC3339, *ti.EndDate)
		if err != nil {
			return t, err
		}
		td.EndDate = endDate
	}
	if ti.Status != nil {
		td.Status = ti.Status.String()
	}
	if ti.CreatedBy != nil {
		td.CreatedBy = *ti.CreatedBy
	}
	if ti.UpdatedBy != nil {
		err = errors.New("Invalid format: UpdatedBy is not null")
		return
	}

	createdTd, err := s.store.Create(&td)
	if err != nil {
		return
	}

	// TODO: create gqlgen custom scalar uint so we do not need to use int()
	t = &model.Todo{
		ID:          strconv.FormatUint(uint64(createdTd.ID), 10),
		ContentCode: createdTd.ContentCode,
		//CreatedDate: createdTd.CreatedAt.UTC().Format(time.RFC3339),
		CreatedDate: createdTd.CreatedAt.Format(time.RFC3339),
		UpdatedDate: createdTd.UpdatedAt.Format(time.RFC3339),
	}
	// FIXME: add nullable fields

	return
}

func (s *todoService) UpdateTodo(id string, ti model.TodoInput) (t *model.Todo, err error) {
	t = &model.Todo{}
	return
}

func NewService(gormdsn string) (TodoService, error) {
	store, err := tododb.NewTodoStore(gormdsn)
	if err != nil {
		return &todoService{}, err
	}
	return &todoService{store: store}, nil
}

// ServiceMiddleware is a chainable behavior modifier for TodoService.
type ServiceMiddleware func(TodoService) TodoService
