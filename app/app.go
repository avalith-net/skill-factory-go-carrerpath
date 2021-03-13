package app

import (
	"log"
	"os"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	if database.ConnectionCheck() == 0 {
		log.Fatal("No connection to the database")
		return
	}
	mapUrls()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	router.Use(cors.Default())
	log.Fatal(router.Run(":" + PORT))
}
