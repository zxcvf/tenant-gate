package tenant_user

import (
	"context"
	"errors"
	"fmt"
	"tenant-gate/internal/entity"
	"tenant-gate/pkg/postgres"

	"github.com/jackc/pgx/v5"
)

type Repo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *Repo {
	return &Repo{pg}
}

func (r *Repo) GetUserTenantRelation(ctx context.Context, userID int64, tenantName string) (entity.TenantUser, error) {
	query := `
		SELECT tu.id, tu.tenant_id, tu.user_id, tu.role_code
		FROM tenants_users tu 
		JOIN tenants t ON t.id = tu.tenant_id 
		WHERE tu.user_id = $1 AND t.tenant_name = $2
	`

	var tu entity.TenantUser
	err := r.Pool.QueryRow(ctx, query, userID, tenantName).Scan(
		&tu.ID,
		&tu.TenantID,
		&tu.UserID,
		&tu.RoleCode,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.TenantUser{}, entity.ErrUserTenantRelationNotFound
		}
		return entity.TenantUser{}, fmt.Errorf("UserRepo - GetUserTenantRelation - r.Pool.QueryRow: %w", err)
	}

	return tu, nil
}
