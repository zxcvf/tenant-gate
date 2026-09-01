package tenant

import (
	"context"
	"tenant-gate/internal/entity"
	"tenant-gate/pkg/postgres"
)

type Repo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *Repo {
	return &Repo{pg}
}

func (r *Repo) Store(ctx context.Context) (string, error) {
	// Implement the logic to store a tenant in the database using r.pg
	return "", nil
}

func (r *Repo) GetTenantByName(ctx context.Context, tenantName string) (entity.Tenant, error) {
	// Implement the logic to retrieve a tenant by name from the database using r.pg
	return entity.Tenant{}, nil
}
