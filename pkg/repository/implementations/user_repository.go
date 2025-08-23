package implementations

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) contracts.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *db.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) CreateBatch(users []db.User) error {
	if len(users) == 0 {
		return nil
	}
	return r.db.Create(&users).Error
}

func (r *userRepository) GetByID(id uuid.UUID) (*db.User, error) {
	var user db.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) GetByEmail(email string) (*db.User, error) {
	var user db.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) GetByUsername(username string) (*db.User, error) {
	var user db.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *userRepository) Update(user *db.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&db.User{}, id).Error
}

func (r *userRepository) GetAll() ([]db.User, error) {
	var users []db.User
	err := r.db.Find(&users).Error
	return users, err
}
