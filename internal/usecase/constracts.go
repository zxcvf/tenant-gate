package usecase

import (
	"context"
	"tenant-gate/internal/entity"
)

type (
	// Tenant -.
	Tenant interface {
		Create(ctx context.Context) (string, error)
		GetTenantByName(ctx context.Context, name string) (entity.Tenant, error)
	}

	// User -.
	User interface {
		Create(ctx context.Context) (string, error)
		Login(ctx context.Context, tenantID, email, password string) (string, error)
		GetUserByID(ctx context.Context, userID string) (string, error)
	}
)

type Manager struct {
	Tenant Tenant
	User   User
}
