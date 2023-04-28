package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GreetBitch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "bitch",
	})
}
