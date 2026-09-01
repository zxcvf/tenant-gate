package entity

type TenantUser struct {
	ID       int   `json:"id" example:"1"`        // primary key
	TenantID int64 `json:"tenant_id" `            // unique together key
	UserID   int64 `json:"user_id"`               // unique together key
	RoleCode int   `json:"role_code" example:"1"` // role code  1=owner, 2=user, etc.
} // @name entity.TenantUser

func (t *TenantUser) getUserRole() (Role, bool) {
	switch t.RoleCode {
	case RoleOwner.RoleCode:
		return RoleOwner, true
	case RoleUser.RoleCode:
		return RoleUser, true
	default:
		return Role{}, false
	}
}
