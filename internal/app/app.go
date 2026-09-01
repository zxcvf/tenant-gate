package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"tenant-gate/config"
	"tenant-gate/internal/restapi"
	"tenant-gate/pkg/httpserver"
	"tenant-gate/pkg/logger"
	"tenant-gate/pkg/postgres"
)

type servers struct {
	http *httpserver.Server

	// Add other servers here (e.g., gRPC, WebSocket, etc.)
}

func (s *servers) startServers() {
	s.http.Start()
}
func (s *servers) shutdownServers(l logger.Interface) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	var err error
	select {
	case sig := <-interrupt:
		l.Info("app - Run - signal: %s", sig.String())
	case err = <-s.http.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	s.shutdownServers(l)
}
func (s *servers) waitForShutdown() {
	s.http.Shutdown()
}

func initServers(l logger.Interface, cfg *config.Config) servers {
	// Initialize the HTTP server
	httpServer := httpserver.New(l,
		httpserver.Port(cfg.Http.Port),
		httpserver.Prefork(cfg.Http.UsePrefork),
	)
	restapi.NewRouter(httpServer.App, cfg, l)

	return servers{http: httpServer}
}

func Run(cfg *config.Config) {
	// Initialize the logger
	l := logger.New(cfg.Log.Level)

	// Initialize the database connection
	pg, err := postgres.New(postgres.BuildUrl(
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Pass,
		cfg.Postgres.Name),
		postgres.MaxPoolSize(cfg.Postgres.PoolMax),
	)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Pool.Close()

	// Initialize and start the HTTP server here
	s := initServers(l, cfg)
	s.startServers()
	s.waitForShutdown()
}
