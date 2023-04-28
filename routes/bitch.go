package routes

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func BitchRoute(r *gin.RouterGroup) {
	r.GET("/", controllers.GreetBitch)
}
