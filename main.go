package main

import (
	routes "fibonacci-spiral-matrix/api/router"
	"fibonacci-spiral-matrix/utils/web"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	port := os.Getenv("API_PORT")
	engine := gin.Default()
	engine.Use(web.AddCosHeaders())
	router := routes.NewRouter(engine)
	router.MapRoutes()
	engine.Run(":" + port)
}
