package controllers

import (
	"fmt"
	"go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateContact(c *gin.Context) {

	var body struct {
		Name    string
		Email   string
		Message string
	}

	c.Bind(&body)
	contact := models.Contact{Name: body.Name, Email: body.Email, Message: body.Message}

	fmt.Println(body)
	c.JSON(http.StatusAccepted, contact)
	// contact := models.Contact{Name: "Bitch"}
	// initializers.DB.Create(&contact)
}
