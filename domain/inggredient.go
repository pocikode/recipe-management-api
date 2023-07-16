package domain

type IngredientContentType string

const (
	IngredientContent IngredientContentType = "content"
	IngredientHeader  IngredientContentType = "header"
)

type Ingredient struct {
	ID       uint64                `gorm:"primaryKey" json:"id"`
	RecipeID uint64                `json:"recipe_id"`
	Type     IngredientContentType `json:"type" validate:"required"`
	Name     string                `json:"name" validate:"required"`
	Sort     uint8                 `json:"sort"`
}
