package jwt

import (
	"strconv"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

// UserClaims represents the JWT claims for a user.
type UserClaims struct {
	UserID int64  `json:"userid,string"`
	Email  string `json:"email"`
	Name   string `json:"name"`

	TenantID   int64  `json:"tenant_id,string"`
	TenantName string `json:"tenant_name"`

	RoleCode int `json:"role"`

	jwtlib.RegisteredClaims
}

func NewUserClaims(userID int64, email, name string, tenantID int64, tenantName string, roleCode int) *UserClaims {
	return &UserClaims{
		UserID:     userID,
		Email:      email,
		Name:       name,
		TenantID:   tenantID,
		TenantName: tenantName,
		RoleCode:   roleCode,

		RegisteredClaims: jwtlib.RegisteredClaims{
			Audience:  []string{"tenant-gate"},
			Subject:   strconv.FormatInt(userID, 10),
			IssuedAt:  jwtlib.NewNumericDate(time.Now()),
			NotBefore: jwtlib.NewNumericDate(time.Now()),
		},
	}
}
