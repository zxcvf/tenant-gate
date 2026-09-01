package entity

import "time"

// User -.
type User struct {
	ID       string `json:"id" example:"user-1"` // primary key
	Username string `json:"username" example:"johndoe"`
	Email    string `json:"email" example:"user@example.com"`
	Phone    string `json:"phone" example:"+1234567890"`
	password string `json:"-" example:"hashed_password"`

	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
} // @name entity.User

type TenantUser struct {
	ID       int    `json:"id" example:"1"`               // primary key
	TenantID string `json:"tenant_id" example:"tenant-1"` // unique key
	UserID   string `json:"user_id" example:"user-1"`     // unique key
	Role     int    `json:"role" example:"1"`             // role code  1=owner, 2=user, etc.
} // @name entity.TenantUser
