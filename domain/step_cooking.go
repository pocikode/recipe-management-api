package domain

type StepContentType string

const (
	StepContent StepContentType = "content"
	StepHeader  StepContentType = "header"
)

type Step struct {
	ID       uint64          `gorm:"primaryKey" json:"id"`
	RecipeID uint64          `json:"recipe_id"`
	Type     StepContentType `json:"type" validate:"required"`
	Text     string          `gorm:"type:text" json:"text" validate:"required"`
	Sort     uint8           `json:"sort"`
}
