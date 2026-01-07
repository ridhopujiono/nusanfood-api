package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhopujiono/nusanfood-api/internal/utils"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "file is required",
		})
		return
	}

	folder := c.DefaultPostForm("folder", "uploads")

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to open file",
		})
		return
	}
	defer src.Close()

	url, err := utils.UploadToCloudinary(
		c.Request.Context(),
		src,
		folder,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "upload failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}
