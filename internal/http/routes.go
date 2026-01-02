package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ridhopujiono/nusanfood-api/internal/modules/food"
)

func Register(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/foods", food.List)
	}
}
