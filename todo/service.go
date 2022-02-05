package todo

import (
	"strconv"
	"time"

	"github.com/siongui/go-graphql-postgresql-todo-example/graph/model"
	"github.com/siongui/go-graphql-postgresql-todo-example/graph/scalar"
	"github.com/siongui/go-graphql-postgresql-todo-example/todo/tododb"
)

type TodoService interface {
	GetTodo(string) (*model.Todo, error)
	TodoPages(model.PaginationInput) (*model.TodoPagination, error)
	TodoSearch(model.TodoSearchInput, model.PaginationInput) (*model.TodoPagination, error)
	CreateTodo(model.CreateTodoInput, string) (*model.Todo, error)
	UpdateTodo(string, model.UpdateTodoInput) (*model.Todo, error)
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

func (s *todoService) CreateTodo(ti model.CreateTodoInput, createdby string) (t *model.Todo, err error) {
	t = &model.Todo{}
	td := tododb.Todo{
		ContentCode: ti.ContentCode,
		ContentName: ti.ContentName,
		Description: ti.Description,
		StartDate:   time.Time(ti.StartDate),
		EndDate:     time.Time(ti.EndDate),
		Status:      ti.Status.String(),
		CreatedBy:   createdby,
	}

	createdTd, err := s.store.Create(&td)
	if err != nil {
		return
	}

	sd := scalar.DateTime(createdTd.StartDate)
	ed := scalar.DateTime(createdTd.EndDate)
	t = &model.Todo{
		ID:          strconv.FormatUint(uint64(createdTd.ID), 10),
		CreatedDate: scalar.DateTime(createdTd.CreatedAt),
		UpdatedDate: scalar.DateTime(createdTd.UpdatedAt),
		ContentCode: createdTd.ContentCode,
		ContentName: &createdTd.ContentName,
		Description: &createdTd.Description,
		StartDate:   &sd,
		EndDate:     &ed,
		Status:      getStatus(createdTd.Status),
		CreatedBy:   &createdTd.CreatedBy,
	}

	return
}

func (s *todoService) UpdateTodo(id string, ti model.UpdateTodoInput) (t *model.Todo, err error) {
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

func getStatus(s string) *model.TodoStatus {
	var v model.TodoStatus
	if s == "Active" {
		v = model.TodoStatusActive
		return &v
	}
	if s == "Inactive" {
		v = model.TodoStatusInactive
		return &v
	}
	return nil
}

// ServiceMiddleware is a chainable behavior modifier for TodoService.
type ServiceMiddleware func(TodoService) TodoService
