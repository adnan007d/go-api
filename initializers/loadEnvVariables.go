package initializers

import (
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	godotenv.Load()

	// if err != nil {
	// 	log.Fatal("Failed to load env variables!")
	// }
}
