package routes

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func AnimeRoute(r *gin.RouterGroup) {
	r.GET("/", controllers.GreetAnime)
}
