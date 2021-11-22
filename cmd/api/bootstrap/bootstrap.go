package bootstrap

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/arkiant/ddd-golang-framework/internal/ping"
	"github.com/arkiant/ddd-golang-framework/kit/cqrs/bus/inmemory"
	"github.com/arkiant/ddd-golang-framework/kit/http/server"
	"github.com/joho/godotenv"
)

// ENVIRONMENT VARIABLES
const (
	ENV = "ENV"
)

const (
	host            = "localhost"
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

	pingService := ping.NewPingService(eventBus)
	pingQueryHandler := ping.NewPingQueryHandler(pingService)
	queryBus.Register(ping.PingQueryType, pingQueryHandler)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, getEnvironment(os.Getenv(ENV)), routes(queryBus, commandBus))
	return srv.Run(ctx)
}

func getEnvironment(string) server.Environment {
	var env server.Environment
	switch os.Getenv(ENV) {
	case string(server.DEVELOPMENT):
		env = server.DEVELOPMENT
	case string(server.STAGING):
		env = server.STAGING
	case string(server.PRODUCTION):
		env = server.PRODUCTION
	}
	return env
}
