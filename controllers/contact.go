package controllers

import (
	"go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateContact(c *gin.Context) {

	var body struct {
		Name    string `form:"name" json:"name" validate:"required"`
		Email   string `form:"email" json:"email" validate:"required,email"`
		Message string `form:"message" json:"message" validate:"required"`
	}

	c.Bind(&body)

	validate := validator.New()

	err := validate.Struct(body)

	if err != nil {
		errorMap := make(map[string]string)
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errorMap[fieldErr.Field()] = fieldErr.Tag()
		}
		c.JSON(http.StatusBadRequest, errorMap)
		return
	}

	contact := models.Contact{Name: body.Name, Email: body.Email, Message: body.Message}

	c.JSON(http.StatusAccepted, contact)
}
