package domain

import (
	"recipe-management/pkg"
	customError "recipe-management/pkg/error"
)

type Category struct {
	ID   uint64 `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

var (
	ErrCategoryNotFound = customError.NewError(404, "category not found.")
	ErrCategoryConflict = customError.NewError(409, "category already exists.")
)

type CategoryUseCase interface {
	Fetch(pagination *pkg.Pagination) (*pkg.Pagination, error)
	Create(c *Category) (*Category, error)
}

type CategoryRepository interface {
	Fetch(pagination pkg.Pagination) (*pkg.Pagination, error)
	Create(c *Category) (*Category, error)
	GetByName(name string) (*Category, error)
	GetByID(id uint64) (*Category, error)
}
