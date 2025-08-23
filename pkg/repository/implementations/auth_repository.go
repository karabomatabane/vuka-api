package implementations

import (
	"gorm.io/gorm"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) contracts.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateAccountCode(accountCode *db.AccountCode) error {
	return r.db.Create(accountCode).Error
}

func (r *authRepository) GetAccountCode() (*db.AccountCode, error) {
	var activeCode *db.AccountCode
	err := r.db.Where("expiration_date > NOW()").First(activeCode).Error
	return activeCode, err
}

func (r *authRepository) DeleteAccountCode(accountCode *db.AccountCode) error {
	return r.db.Delete(accountCode).Error
}
