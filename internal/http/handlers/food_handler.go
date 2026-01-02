package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ridhopujiono/nusanfood-api/internal/modules/food"
)

func GetFoods(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	withNutrition := c.Query("with_nutrition") == "true"

	foods, total := food.GetFoods(page, perPage, withNutrition)

	c.JSON(http.StatusOK, gin.H{
		"data": foods,
		"meta": gin.H{
			"page":     page,
			"per_page": perPage,
			"total":    total,
		},
	})
}
