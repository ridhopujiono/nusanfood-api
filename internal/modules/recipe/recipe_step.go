package recipe

type RecipeStep struct {
	ID          uint `gorm:"primaryKey"`
	RecipeID    uint
	Description string
	PhotoURL    string
	StepOrder   int
}

func (RecipeStep) TableName() string {
	return "recipe_steps"
}
