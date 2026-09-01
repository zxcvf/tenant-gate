package user

import (
	"context"
	repo "tenant-gate/internal/repo"
	"tenant-gate/pkg/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UseCase struct {
	repo         repo.User
	relationRepo repo.TenantUserRelation
	jwt          *jwt.Manager
}

func New(r repo.User, t repo.TenantUserRelation, j *jwt.Manager) *UseCase {
	return &UseCase{
		repo:         r,
		relationRepo: t,
		jwt:          j,
	}
}

func (u *UseCase) Create(ctx context.Context) (string, error) {
	return "", nil
}

func (u *UseCase) Login(ctx context.Context, tenantName, email, password string) (string, error) {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	// check password
	if !user.CheckPassword(password) {
		return "", fiber.ErrUnauthorized
	}

	// check user-tenant relation
	tn, err := u.relationRepo.GetUserTenantRelation(ctx, user.ID, tenantName)
	if err != nil {
		return "", err
	}

	// todo: generate jwt token
	log.Debug(tn)
	return "token", nil
}

func (u *UseCase) GetUserByID(ctx context.Context, userID string) (string, error) {
	return "", nil
}
