package entity

import (
	"time"
)

// Tenant -.
type Tenant struct {
	ID         int64     `json:"id,string"` // primary key
	TenantName string    `json:"tenant_name" example:"Tenant Name"`
	Email      string    `json:"email" example:"tenant@example.com"`
	CreatedBy  string    `json:"created_by" example:"user-1"`
	CreatedAt  time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt  time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
} // @name entity.Tenant
