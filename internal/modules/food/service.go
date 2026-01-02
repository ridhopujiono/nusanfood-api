package food

func ListFoods(page, perPage int, withNutrition bool) (any) {
	foods, total := GetFoods(page, perPage, withNutrition)

	return map[string]any{
		"data": foods,
		"meta": map[string]any{
			"page": page,
			"per_page": perPage,
			"total": total,
		},
	}
}
