package todo

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/siongui/go-graphql-postgresql-todo-example/graph/model"
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
			"output", *t,
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
			"input", pi,
			"output", *tp,
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
			"method", "TodoPages",
			"input(model.TodoSearchInput)", tsi,
			"input(model.PaginationInput)", pi,
			"output", *tp,
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
			"method", "TodoPages",
			"input", ti,
			"createdby", createdby,
			"output", *t,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	t, err = mw.next.CreateTodo(ti, createdby)
	return
}

func (mw *loggingMiddleware) UpdateTodo(id string, ti model.UpdateTodoInput) (t *model.Todo, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "TodoPages",
			"input(id)", id,
			"input(model.TodoInput)", ti,
			"output", *t,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	t, err = mw.next.UpdateTodo(id, ti)
	return
}
