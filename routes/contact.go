package routes

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func ContactRouter(r *gin.RouterGroup) {
	r.POST("/", controllers.ContactCreate)
}
