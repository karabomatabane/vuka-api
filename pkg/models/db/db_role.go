package db

import "vuka-api/pkg/models/db/role"

type Role struct {
	Model
	Name                   string                  `json:"name"`
	RoleSectionPermissions []RoleSectionPermission `gorm:"foreignKey:RoleID"`
}

func (r *Role) ToDomain() role.Response {
	permissionsMap := make(map[string][]string)

	for _, p := range r.RoleSectionPermissions {
		sectionName := p.Section.Name
		permissionName := p.Permission.Name
		permissionsMap[sectionName] = append(permissionsMap[sectionName], permissionName)
	}

	return role.Response{
		ID:          r.ID,
		Name:        r.Name,
		Permissions: permissionsMap,
	}
}
