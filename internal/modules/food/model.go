package food

// =======================
// FOODS
// =======================

type Food struct {
	ID       uint          `json:"id" gorm:"column:id;primaryKey"`
	Name     string        `json:"name" gorm:"column:name"`
	FoodType string        `json:"food_type" gorm:"column:food_type"`
	Servings []FoodServing `json:"servings" gorm:"foreignKey:FoodID;references:ID"`
}

func (Food) TableName() string {
	return "foods"
}

// =======================
// FOOD SERVINGS
// =======================

type FoodServing struct {
	ID           uint    `json:"id" gorm:"column:id;primaryKey"`
	FoodID       uint    `json:"food_id" gorm:"column:food_id"`
	ServingLabel string  `json:"serving_label" gorm:"column:serving_label"`
	MetricAmount float64 `json:"metric_amount" gorm:"column:metric_amount"`
	MetricUnit   string  `json:"metric_unit" gorm:"column:metric_unit"`

	Nutrition NutritionFact `json:"nutrition" gorm:"foreignKey:ServingID;references:ID"`
}

func (FoodServing) TableName() string {
	return "food_servings"
}

// =======================
// NUTRITION FACTS  ✅
// =======================

type NutritionFact struct {
	ID        uint `json:"id" gorm:"column:id;primaryKey"`
	ServingID uint `json:"serving_id" gorm:"column:serving_id"`

	Calories     float64 `json:"calories" gorm:"column:calories"`
	Carbohydrate float64 `json:"carbohydrate" gorm:"column:carbohydrate"`
	Protein      float64 `json:"protein" gorm:"column:protein"`
	Fat          float64 `json:"fat" gorm:"column:fat"`

	SaturatedFat       float64 `json:"saturated_fat" gorm:"column:saturated_fat"`
	PolyunsaturatedFat float64 `json:"polyunsaturated_fat" gorm:"column:polyunsaturated_fat"`
	MonounsaturatedFat float64 `json:"monounsaturated_fat" gorm:"column:monounsaturated_fat"`

	Fiber float64 `json:"fiber" gorm:"column:fiber"`
	Sugar float64 `json:"sugar" gorm:"column:sugar"`

	Cholesterol float64 `json:"cholesterol" gorm:"column:cholesterol"`
	Sodium      float64 `json:"sodium" gorm:"column:sodium"`
	Potassium   float64 `json:"potassium" gorm:"column:potassium"`

	VitaminA float64 `json:"vitamin_a" gorm:"column:vitamin_a"`
	VitaminC float64 `json:"vitamin_c" gorm:"column:vitamin_c"`
	Calcium  float64 `json:"calcium" gorm:"column:calcium"`
	Iron     float64 `json:"iron" gorm:"column:iron"`
}

func (NutritionFact) TableName() string {
	return "nutrition_facts"
}
