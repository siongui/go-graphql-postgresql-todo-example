package main

import (
	"context"
	"os"

	"github.com/siongui/go-kit-gqlgen-postgres-todo-example/config"
	"github.com/siongui/go-kit-gqlgen-postgres-todo-example/graph"
	"github.com/siongui/go-kit-gqlgen-postgres-todo-example/graph/generated"
	"github.com/siongui/go-kit-gqlgen-postgres-todo-example/todo"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// set up middleware to pass *gin.Context to context.Context in resolver
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), graph.GetCtxKey(), c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Load config
	if err := config.LoadConfig(); err != nil {
		logger.Log("err", err.Error())
		os.Exit(1)
	}
	logger.Log("msg", "App config")
	logger.Log("port", config.Config.App.Port)
	logger.Log("timezone", config.Config.App.Timezone)
	gormdsn := "host=" + config.Config.Database.Postgres.Host +
		" user=" + config.Config.Database.Postgres.User +
		" password=" + config.Config.Database.Postgres.Password +
		" dbname=" + config.Config.Database.Postgres.Dbname +
		" port=" + config.Config.Database.Postgres.Port +
		" sslmode=disable TimeZone=" + config.Config.App.Timezone
	logger.Log("gorm_dsn", gormdsn)

	// set up prometheus
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

	// set up todo service and middleware
	var svc todo.TodoService
	svc, err := todo.NewService(gormdsn)
	if err != nil {
		logger.Log("err", err.Error())
		os.Exit(1)
	}
	svc = todo.NewLoggingMiddleware(logger, svc)
	svc = todo.NewInstrumentingMiddleware(requestCount, requestLatency, svc)
	eps := todo.MakeEndPoints(svc, logger)

	// Make GraphQL handler
	gqlhdr := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			Logger:             logger,
			GetTodoEndpoint:    eps.GetTodoEndpoint,
			TodoPagesEndpoint:  eps.TodoPagesEndpoint,
			TodoSearchEndpoint: eps.TodoSearchEndpoint,
			CreateTodoEndpoint: eps.CreateTodoEndpoint,
			UpdateTodoEndpoint: eps.UpdateTodoEndpoint,
		},
	}))

	router := gin.New()
	router.Use(GinContextToContextMiddleware())
	router.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/query")))
	router.POST("/query", gin.WrapH(gqlhdr))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	logger.Log("msg", "connect to http://localhost:"+config.Config.App.Port+"/ for GraphQL playground")
	logger.Log("err", router.Run(":"+config.Config.App.Port))
}
