package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GreetFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from the file controller!",
	})
}

func SingleFileUpload(c *gin.Context) {

	formData, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error getting form data",
		})
		return
	}

	name := formData.Value["name"]
	message := formData.Value["message"]

	// Check if the file is present
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No file uploaded",
		})
		return
	}

	fmt.Println(file.Size, file.Filename)

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from the file controller!",
		"data":    gin.H{"name": name, "message": message},
	})
}
