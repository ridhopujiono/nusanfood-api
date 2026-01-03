package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ridhopujiono/nusanfood-api/internal/config"
	"github.com/ridhopujiono/nusanfood-api/internal/database"
	"github.com/ridhopujiono/nusanfood-api/internal/http/handlers"
	"github.com/ridhopujiono/nusanfood-api/internal/http/middleware"
)

func main() {
	config.Load()
	database.Connect()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	protected := api.Group("/")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/foods", handlers.GetFoods)
	}

	port := config.Get("APP_PORT", "8080")
	r.Run(":" + port)

}
