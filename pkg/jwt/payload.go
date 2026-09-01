package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims represents the JWT claims for a user.
type CustomClaims struct {
	UserID   int64 `json:"sub, string"`
	TenantID int64 `json:"tenant_id, string"`
	RoleCode int   `json:"role"`
	jwt.RegisteredClaims
}
