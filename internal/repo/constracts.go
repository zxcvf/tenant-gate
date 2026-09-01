package usecase

import (
	"context"
	"tenant-gate/internal/entity"
)

type (
	// Tenant -.
	Tenant interface {
		Store(ctx context.Context) (string, error)
		GetTenantByName(ctx context.Context, tenantName string) (entity.Tenant, error)
	}

	// User -.
	User interface {
		Store(ctx context.Context) (string, error)
		GetUserByID(ctx context.Context, userID int64) (string, error)
		GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	}

	TenantUserRelation interface {
		GetUserTenantRelation(ctx context.Context, userID int64, tenantName string) (entity.TenantUser, error)
	}
)
