package recipe

type RecipeIngredient struct {
	ID            uint `gorm:"primaryKey"`
	RecipeID      uint
	FoodID        uint
	FoodServingID uint
	Quantity      float64
}

func (RecipeIngredient) TableName() string {
	return "recipe_ingredients"
}
