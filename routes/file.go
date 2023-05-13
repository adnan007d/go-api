package routes

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func FileRoute(r *gin.RouterGroup) {
	r.GET("/", controllers.GreetFile)
	r.POST("/single", controllers.SingleFileUpload)
}
