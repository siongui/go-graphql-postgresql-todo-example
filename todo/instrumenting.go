package todo

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/siongui/go-graphql-postgresql-todo-example/graph/model"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           TodoService
}

func NewInstrumentingMiddleware(
	counter metrics.Counter,
	latency metrics.Histogram,
	svc TodoService) TodoService {

	return &instrumentingMiddleware{
		requestCount:   counter,
		requestLatency: latency,
		next:           svc,
	}
}

func (mw *instrumentingMiddleware) GetTodo(id string) (t *model.Todo, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetTodo", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	t, err = mw.next.GetTodo(id)
	return
}

func (mw *instrumentingMiddleware) TodoPages(pi model.PaginationInput) (tp *model.TodoPagination, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "TodoPages", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	tp, err = mw.next.TodoPages(pi)
	return
}

func (mw *instrumentingMiddleware) TodoSearch(tsi model.TodoSearchInput, pi model.PaginationInput) (tp *model.TodoPagination, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "TodoSearch", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	tp, err = mw.next.TodoSearch(tsi, pi)
	return
}

func (mw *instrumentingMiddleware) CreateTodo(ti model.TodoInput) (t *model.Todo, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateTodo", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	t, err = mw.next.CreateTodo(ti)
	return
}

func (mw *instrumentingMiddleware) UpdateTodo(id string, ti model.TodoInput) (t *model.Todo, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateTodo", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	t, err = mw.next.UpdateTodo(id, ti)
	return
}
