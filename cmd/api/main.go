package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ridhopujiono/nusanfood-api/internal/config"
	"github.com/ridhopujiono/nusanfood-api/internal/database"
	httpRoutes "github.com/ridhopujiono/nusanfood-api/internal/http"
)

func main() {
	config.Load()
	database.Connect()

	r := gin.Default()
	httpRoutes.Register(r)

    port := config.Get("APP_PORT", "8080")
    r.Run(":" + port)

}
