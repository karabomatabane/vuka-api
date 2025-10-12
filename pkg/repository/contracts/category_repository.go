package contracts

import "vuka-api/pkg/models/db"

type CategoryRepository interface {
	FindAll() ([]db.Category, error)
	FindIn(field string, values []any, target any) error
	FindByName(name string) (*db.Category, error)
	Create(category *db.Category) error
}