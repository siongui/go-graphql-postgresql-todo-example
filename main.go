package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/siongui/go-graphql-postgresql-todo-example/todo"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "todo_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "todo_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	var svc todo.TodoService
	svc = todo.NewService()
	svc = todo.NewLoggingMiddleware(logger, svc)
	svc = todo.NewInstrumentingMiddleware(requestCount, requestLatency, svc)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", todo.MakeGraphQLHandler(svc, logger))
	http.Handle("/metrics", promhttp.Handler())

	logger.Log("msg", "connect to http://localhost:"+port+"/ for GraphQL playground")
	logger.Log("err", http.ListenAndServe(":"+port, nil))
}
