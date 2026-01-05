package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhopujiono/nusanfood-api/internal/modules/recipe"
)

func CreateRecipe(c *gin.Context) {
	var req recipe.CreateRecipeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"errors":  err.Error(),
		})
		return
	}

	userID := c.GetUint("user_id") // dari JWT middleware

	data, err := recipe.CreateRecipe(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create recipe",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Recipe created",
		"data":    data,
	})
}
