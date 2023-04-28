package main

import (
	"go-api/initializers"
	"go-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Contact{})
}
