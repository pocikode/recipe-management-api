package http

import "time"

type RecipeListSerializer struct {
	ID           uint64     `json:"id"`
	CategoryID   uint64     `json:"category_id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	CookingTime  uint8      `json:"cooking_time"`
	ThumbnailURL string     `json:"thumbnail_url"`
	VideoURL     string     `json:"video_url"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
