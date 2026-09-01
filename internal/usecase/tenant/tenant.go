package tenant

import (
	"context"
	"tenant-gate/internal/entity"
	repo "tenant-gate/internal/repo"
)

type UseCase struct {
	repo repo.Tenant
}

func New(r repo.Tenant) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (u *UseCase) Create(ctx context.Context) (string, error) {
	return "", nil
}

func (u *UseCase) GetTenantByName(ctx context.Context, name string) (entity.Tenant, error) {
	return entity.Tenant{}, nil
}
