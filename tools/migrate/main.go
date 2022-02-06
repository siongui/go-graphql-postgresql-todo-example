package main

import (
	"flag"
	"os"

	"github.com/go-kit/log"
	"github.com/siongui/go-kit-gqlgen-postgres-todo-example/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	isSSL := flag.Bool("ssl", true, "is SSL mode?")
	dir := flag.String("dir", "migrations", "path to migrations dir")
	flag.Parse()

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

	dsn := "postgres://" +
		config.Config.Database.Postgres.User + ":" +
		config.Config.Database.Postgres.Password + "@" +
		config.Config.Database.Postgres.Host + ":" +
		config.Config.Database.Postgres.Port + "/" +
		config.Config.Database.Postgres.Dbname

	if !*isSSL {
		dsn += "?sslmode=disable"
	}

	logger.Log("DSN", dsn)

	logger.Log("migrations_dir", *dir)
	m, err := migrate.New("file://"+*dir, dsn)
	if err != nil {
		logger.Log("err", err.Error())
		os.Exit(1)
	}
	if err = m.Up(); err != nil {
		logger.Log("err", err.Error())
		os.Exit(1)
	}
}
