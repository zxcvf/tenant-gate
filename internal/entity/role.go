package entity

type Role struct {
	RoleName string
	RoleCode int
}

// Role -.
var (
	RoleOwner = Role{
		RoleName: "Owner",
		RoleCode: 1,
	}

	RoleUser = Role{
		RoleName: "User",
		RoleCode: 2,
	}
)

func (r Role) GetRoleName() string {
	return r.RoleName
}

func (r Role) GetRoleCode() int {
	return r.RoleCode
}

func GetRoleByCode(code int) (Role, bool) {
	switch code {
	case RoleOwner.RoleCode:
		return RoleOwner, true
	case RoleUser.RoleCode:
		return RoleUser, true
	default:
		return Role{}, false
	}
}
