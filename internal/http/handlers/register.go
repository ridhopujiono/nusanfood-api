package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhopujiono/nusanfood-api/internal/auth"
	"github.com/ridhopujiono/nusanfood-api/internal/modules/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// cek email sudah ada
	if _, err := user.FindByEmail(req.Email); err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Email already registered",
		})
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash password",
		})
		return
	}

	newUser := &user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := user.Create(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	token, err := auth.GenerateUserToken(newUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"access_token": token,
		"token_type":   "Bearer",
		"user": gin.H{
			"id":    newUser.ID,
			"email": newUser.Email,
		},
	})
}
