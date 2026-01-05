package recipe

import "time"

// =======================
// RECIPES
// =======================

type Recipe struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	UserID          uint   `gorm:"column:user_id" json:"user_id"`
	Title           string `gorm:"column:title" json:"title"`
	CoverUrl        string `gorm:"column:cover_url" json:"cover_url"`
	Description     string `gorm:"column:description" json:"description"`
	CookTimeMinutes int    `gorm:"column:cook_time_minutes" json:"cook_time_minutes"`
	Servings        int    `gorm:"column:servings" json:"servings"`

	Ingredients []RecipeIngredient `gorm:"foreignKey:RecipeID"`
	Steps       []RecipeStep       `gorm:"foreignKey:RecipeID"`

	CreatedAt time.Time
}

func (Recipe) TableName() string {
	return "recipes"
}
