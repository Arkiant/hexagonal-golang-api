package bootstrap

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/arkiant/hexagonal-golang-api/internal/sending/ping"
	"github.com/arkiant/hexagonal-golang-api/kit/cqrs/bus/inmemory"
	"github.com/arkiant/hexagonal-golang-api/kit/http/server"
	"github.com/joho/godotenv"
)

// ENVIRONMENT VARIABLES
const (
	ENV = "ENV"
)

const (
	host            = "0.0.0.0"
	port            = 8080
	shutdownTimeout = 10 * time.Second
	dbTimeout       = 5 * time.Second
)

func Run() error {

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
		queryBus   = inmemory.NewQueryBus()
	)

	var (
		_, base, _, _   = runtime.Caller(0)
		basePath        = filepath.Dir(base)
		environmentPath = filepath.Join(basePath, "../../../", ".env")
	)

	err := godotenv.Load(environmentPath)
	if err != nil {
		return err
	}

	_ = eventBus

	queryBus.Register(ping.QueryType, ping.NewQueryHandler(ping.NewService()))

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, getEnvironment(os.Getenv(ENV)), routes(queryBus, commandBus))
	return srv.Run(ctx)
}

func getEnvironment(environment string) server.Environment {
	var env server.Environment
	switch environment {
	case string(server.DEVELOPMENT):
		env = server.DEVELOPMENT
	case string(server.STAGING):
		env = server.STAGING
	case string(server.PRODUCTION):
		env = server.PRODUCTION
	}
	return env
}
