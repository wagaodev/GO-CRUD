package main

import (
	"api/services/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Falha ao carregar o arquivo .env")
	}

	app := gin.Default()

	routes.AppRoutes(app)

	// get the environment variable of APP_HOST
	appHost := os.Getenv("APP_HOST")
	if appHost == "" {
		panic("Falha ao obter a variável de ambiente APP_HOST")
	}

	app.Run(appHost)
}
