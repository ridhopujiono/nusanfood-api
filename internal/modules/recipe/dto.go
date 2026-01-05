package recipe

type CreateRecipeRequest struct {
	Title           string `json:"title" binding:"required"`
	CoverUrl        string `json:"cover_url"`
	Description     string `json:"description"`
	CookTimeMinutes int    `json:"cook_time_minutes"`
	Servings        int    `json:"servings"`

	Ingredients []struct {
		FoodID        uint    `json:"food_id" binding:"required"`
		FoodServingID uint    `json:"food_serving_id" binding:"required"`
		Quantity      float64 `json:"quantity" binding:"required"`
	} `json:"ingredients" binding:"required"`

	Steps []struct {
		Description string `json:"description" binding:"required"`
		PhotoURL    string `json:"photo_url"`
	} `json:"steps" binding:"required"`
}
