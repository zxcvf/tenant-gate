package entity

import "errors"

var (
	ErrTenantNotFound = errors.New("tenant not found")
	ErrUserNotFound   = errors.New("user not found")
	ErrUserExists     = errors.New("user already exists")

	ErrUserTenantRelationNotFound = errors.New("user-tenant relation not found")
)
