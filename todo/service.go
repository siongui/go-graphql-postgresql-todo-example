package todo

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/siongui/go-kit-gqlgen-postgres-todo-example/graph/model"
	"github.com/siongui/go-kit-gqlgen-postgres-todo-example/graph/scalar"
	"github.com/siongui/go-kit-gqlgen-postgres-todo-example/todo/tododb"
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
	td, err := s.store.GetTodo(id)
	if err != nil {
		return
	}
	t = toModelTodo(td)
	return
}

func (s *todoService) TodoPages(pi model.PaginationInput) (tp *model.TodoPagination, err error) {
	tp = &model.TodoPagination{}
	// record count per page
	count := pi.Count
	// n-th page
	page := pi.Page

	if count < 1 {
		err = errors.New("TodoPages: count must >= 1")
		return
	}

	if page < 1 {
		err = errors.New("TodoPages: page must >= 1")
		return
	}

	todos, totalCount, err := s.store.Pages(count, page)
	if err != nil {
		return
	}

	var modelTodos []*model.Todo
	for _, todo := range todos {
		modelTodos = append(modelTodos, toModelTodo(todo))
	}

	tp = &model.TodoPagination{
		PaginationInfo: &model.PaginationInfo{
			TotalCount:  int(totalCount),
			CurrentPage: page,
			TotalPages:  int(math.Ceil(float64(totalCount) / float64(count))),
		},
		Todos: modelTodos,
	}

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

	t = toModelTodo(*createdTd)

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
	if s == model.TodoStatusActive.String() {
		v = model.TodoStatusActive
		return &v
	}
	if s == model.TodoStatusInactive.String() {
		v = model.TodoStatusInactive
		return &v
	}
	return nil
}

func toModelTodo(todo tododb.Todo) *model.Todo {
	sd := scalar.DateTime(todo.StartDate)
	ed := scalar.DateTime(todo.EndDate)
	mtd := model.Todo{
		ID:          strconv.FormatUint(uint64(todo.ID), 10),
		CreatedDate: scalar.DateTime(todo.CreatedAt),
		UpdatedDate: scalar.DateTime(todo.UpdatedAt),
		ContentCode: todo.ContentCode,
		ContentName: &todo.ContentName,
		Description: &todo.Description,
		StartDate:   &sd,
		EndDate:     &ed,
		Status:      getStatus(todo.Status),
		CreatedBy:   &todo.CreatedBy,
		UpdatedBy:   &todo.UpdatedBy,
	}

	if todo.UpdatedBy == "" {
		mtd.UpdatedBy = nil
	}

	return &mtd
}

// ServiceMiddleware is a chainable behavior modifier for TodoService.
type ServiceMiddleware func(TodoService) TodoService
