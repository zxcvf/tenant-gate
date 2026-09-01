package request

// Login -.
type Login struct {
	TenantName string `json:"tenant_name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
} // @name v1.Login
