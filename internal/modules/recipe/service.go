package recipe

import (
	"errors"

	"github.com/ridhopujiono/nusanfood-api/internal/database"
)

func CreateRecipe(userID uint, req CreateRecipeRequest) (*Recipe, error) {
	tx := database.DB.Begin()

	if tx.Error != nil {
		return nil, tx.Error
	}

	recipe := Recipe{
		UserID:          userID,
		Title:           req.Title,
		CoverUrl:        req.CoverUrl,
		Description:     req.Description,
		CookTimeMinutes: req.CookTimeMinutes,
		Servings:        req.Servings,
	}

	if err := tx.Create(&recipe).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// ingredients
	for _, ing := range req.Ingredients {
		item := RecipeIngredient{
			RecipeID:      recipe.ID,
			FoodID:        ing.FoodID,
			FoodServingID: ing.FoodServingID,
			Quantity:      ing.Quantity,
		}
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// steps
	for i, step := range req.Steps {
		item := RecipeStep{
			RecipeID:    recipe.ID,
			Description: step.Description,
			PhotoURL:    step.PhotoURL,
			StepOrder:   i + 1,
		}
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errors.New("failed to commit transaction")
	}

	return &recipe, nil
}
