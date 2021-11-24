package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type Environment string

const (
	DEVELOPMENT Environment = "dev"
	STAGING     Environment = "qa"
	PRODUCTION  Environment = "prod"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	shutdownTimeout time.Duration

	routes []Route
}

type Route struct {
	method   string
	endpoint string
	handler  gin.HandlerFunc
}

func NewRoute(method, endpoint string, handler gin.HandlerFunc) Route {
	return Route{method: method, endpoint: endpoint, handler: handler}
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, environment Environment, routes []Route) (context.Context, Server) {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		shutdownTimeout: shutdownTimeout,

		routes: routes,
	}

	srv.engine.Use(gin.Recovery(), gin.Logger())
	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) registerRoutes() {
	if !s.hasRoutes() {
		panic("no routes found")
	}

	for _, route := range s.routes {
		s.engine.Handle(route.method, route.endpoint, route.handler)
	}
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shutdown", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func (s *Server) hasRoutes() bool {
	return len(s.routes) > 0
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
