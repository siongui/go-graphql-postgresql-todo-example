package todo

import (
	"fmt"
	"time"

	"github.com/go-kit/log"
	"github.com/siongui/go-kit-gqlgen-postgres-todo-example/graph/model"
)

type loggingMiddleware struct {
	logger log.Logger
	next   TodoService
}

func NewLoggingMiddleware(logger log.Logger, s TodoService) TodoService {
	return &loggingMiddleware{logger, s}
}

func (mw *loggingMiddleware) GetTodo(id string) (t *model.Todo, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetTodo",
			"input", id,
			"output", fmt.Sprintf("%#v", t),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	t, err = mw.next.GetTodo(id)
	return
}

func (mw *loggingMiddleware) TodoPages(pi model.PaginationInput) (tp *model.TodoPagination, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "TodoPages",
			"input", fmt.Sprintf("%#v", pi),
			"output", fmt.Sprintf("%#v", tp),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	tp, err = mw.next.TodoPages(pi)
	return
}

func (mw *loggingMiddleware) TodoSearch(tsi model.TodoSearchInput, pi model.PaginationInput) (tp *model.TodoPagination, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "TodoSearch",
			"input(model.TodoSearchInput)", fmt.Sprintf("%#v", tsi),
			"input(model.PaginationInput)", fmt.Sprintf("%#v", pi),
			"output", fmt.Sprintf("%#v", tp),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	tp, err = mw.next.TodoSearch(tsi, pi)
	return
}

func (mw *loggingMiddleware) CreateTodo(ti model.CreateTodoInput, createdby string) (t *model.Todo, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "CreateTodo",
			// https://stackoverflow.com/questions/16331063/how-can-i-get-the-string-representation-of-a-struct
			"input", fmt.Sprintf("%#v", ti),
			"createdby", createdby,
			"output", fmt.Sprintf("%#v", t),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	t, err = mw.next.CreateTodo(ti, createdby)
	return
}

func (mw *loggingMiddleware) UpdateTodo(id string, ti model.UpdateTodoInput, updatedby string) (t *model.Todo, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "UpdateTodo",
			"input(id)", id,
			"input(model.TodoInput)", fmt.Sprintf("%#v", ti),
			"input(updatedby)", updatedby,
			"output", fmt.Sprintf("%#v", t),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	t, err = mw.next.UpdateTodo(id, ti, updatedby)
	return
}
