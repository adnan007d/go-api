package handler

import (
	"go-api/initializers"
	"go-api/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Nothing to see here, move along.",
	})
}

func init() {
	app = gin.New()

	initializers.LoadEnvVariables()
	initializers.ConnectDB()

	app.MaxMultipartMemory = 3 << 20 // 3 MiB

	app.GET("/", Home)
	animeRouter := app.Group("/api/anime")
	contactRouter := app.Group("/api/contact")
	fileRouter := app.Group("/api/file")

	routes.AnimeRoute(animeRouter)
	routes.ContactRouter(contactRouter)
	routes.FileRoute(fileRouter)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
