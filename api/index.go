package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func BitchRoute(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "bitch",
		})
	})
}

func AnimeRoute(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "anime",
		})
	})
}

func init() {
	app = gin.New()

	bitchRouter := app.Group("/api/bitch")
	animeRouter := app.Group("/api/anime")
	BitchRoute(bitchRouter)
	AnimeRoute(animeRouter)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
