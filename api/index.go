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

func init() {
	app = gin.New()

	initializers.LoadEnvVariables()
	initializers.ConnectDB()

	bitchRouter := app.Group("/api/bitch")
	animeRouter := app.Group("/api/anime")
	contactRouter := app.Group("/api/contact")
	routes.BitchRoute(bitchRouter)
	routes.AnimeRoute(animeRouter)
	routes.ContactRouter(contactRouter)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
