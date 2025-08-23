package contracts

import (
	"vuka-api/pkg/models/db"
)

type AuthRepository interface {
	CreateAccountCode(accountCode *db.AccountCode) error
	GetAccountCode() (*db.AccountCode, error)
	DeleteAccountCode(accountCode *db.AccountCode) error
}
