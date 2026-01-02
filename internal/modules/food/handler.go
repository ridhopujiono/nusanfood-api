package food

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	withNutrition := c.Query("with_nutrition") == "true"

	res := ListFoods(page, perPage, withNutrition)

	c.JSON(200, res)
}
