package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"tenant-gate/config"
	"tenant-gate/internal/controller/restapi"
	"tenant-gate/internal/usecase"
	"tenant-gate/pkg/httpserver"
	"tenant-gate/pkg/jwt"
	"tenant-gate/pkg/logger"
	"tenant-gate/pkg/postgres"

	tenantRepo "tenant-gate/internal/repo/persistent/tenant"
	tenantUserRepo "tenant-gate/internal/repo/persistent/tenant_user"

	userRepo "tenant-gate/internal/repo/persistent/user"
	tenantUsecase "tenant-gate/internal/usecase/tenant"
	userUsecase "tenant-gate/internal/usecase/user"
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

func initUseCases(pg *postgres.Postgres, jwtManager *jwt.Manager) *usecase.Manager {
	// Initialize repositories
	return &usecase.Manager{
		Tenant: tenantUsecase.New(tenantRepo.New(pg)),
		User:   userUsecase.New(userRepo.New(pg), tenantUserRepo.New(pg), jwtManager),
	}
}

func initServers(l logger.Interface, cfg *config.Config, usecase *usecase.Manager, jwtManager *jwt.Manager) servers {
	// Initialize the HTTP server
	httpServer := httpserver.New(l,
		httpserver.Port(cfg.Http.Port),
		httpserver.Prefork(cfg.Http.UsePrefork),
	)
	restapi.NewRouter(httpServer.App, cfg, usecase, jwtManager, l)

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

	// Initialize JWT manager
	jwtManager := jwt.New(cfg.Jwt.Secret, cfg.Jwt.TokenExpiry)

	// Initialize use cases
	usecases := initUseCases(pg, jwtManager)

	// Initialize and start the HTTP server here
	s := initServers(l, cfg, usecases, jwtManager)
	s.startServers()
	s.waitForShutdown()
}
