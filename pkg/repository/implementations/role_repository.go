package implementations

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) contracts.RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(role *db.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) Update(role *db.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) GetById(id string) (*db.Role, error) {
	var role db.Role
	err := r.db.First(&role, id).Error
	return &role, err
}

func (r *roleRepository) GetWithPermissions(id string) (*db.Role, error) {
	var role db.Role
	err := r.db.Preload("Permissions").First(&role, id).Error
	return &role, err
}

func (r *roleRepository) GetAll() ([]db.Role, error) {
	var roles []db.Role
	err := r.db.Find(&roles).Error
	return roles, err
}

func (r *roleRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&db.Role{}, id).Error
}
