package domain

import (
	"recipe-management/pkg"
	error2 "recipe-management/pkg/error"
	"time"
)

var (
	ErrRecipeNotFound = error2.NewError(404, "recipe not found.")
)

type Recipe struct {
	ID           uint64       `gorm:"primaryKey" json:"id"`
	CategoryID   uint64       `json:"category_id" validate:"required"`
	Title        string       `json:"title" validate:"required"`
	Description  string       `json:"description"`
	CookingTime  uint8        `json:"cooking_time" validate:"required"`
	ThumbnailURL string       `json:"thumbnail_url" validate:"required"`
	VideoURL     string       `json:"video_url"`
	CreatedAt    *time.Time   `json:"created_at"`
	UpdatedAt    *time.Time   `json:"updated_at"`
	Category     Category     `json:"category"`
	Ingredients  []Ingredient `json:"ingredients" validate:"required" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Steps        []Step       `json:"steps" validate:"required" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type RecipeUseCase interface {
	Fetch(pagination *pkg.Pagination) error
	GetByID(id uint64) (*Recipe, error)
	Create(recipe *Recipe) (*Recipe, error)
	Update(id uint64, updateData *Recipe) (*Recipe, error)
	Delete(id uint64) error
}

type RecipeRepository interface {
	Fetch(pagination *pkg.Pagination) error
	GetByID(id uint64) (*Recipe, error)
	Create(recipe *Recipe) (*Recipe, error)
	Update(recipe *Recipe, updateData *Recipe) (*Recipe, error)
	Delete(recipe *Recipe) error
}
