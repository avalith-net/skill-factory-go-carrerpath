package app

import (
	"log"

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
	router.Use(cors.Default())
	log.Fatal(router.Run())
}
