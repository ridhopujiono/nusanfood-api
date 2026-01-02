package food

import (
	"github.com/ridhopujiono/nusanfood-api/internal/database"
	"gorm.io/gorm"
)

func GetFoods(page, perPage int, withNutrition bool) ([]Food, int64) {
	var foods []Food
	var total int64

	q := database.DB.Model(&Food{})

	q.Count(&total)

	if withNutrition {
		q = q.
			Preload("Servings", func(db *gorm.DB) *gorm.DB {
				return db.Order("id ASC")
			}).
			Preload("Servings.Nutrition")
	}

	q.Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&foods)

	return foods, total
}
