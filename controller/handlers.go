package controller

import (
	"log"
	"os"

	"github.com/avalith-net/skill-factory-go-carrerpath/middlewares"
	"github.com/avalith-net/skill-factory-go-carrerpath/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//LaunchingServer seteo mi puerto, el handler y pongo a escuchar al servidor
func LaunchingServer() {
	router := gin.Default()
	router.Use(middlewares.CheckDB())
	router.POST("/register", routers.Register)
	router.POST("/login", routers.Login)
	router.GET("/userpath", middlewares.ValidateJWT(), routers.ShowUserPath)
	router.POST("/passwordRecovery", routers.PasswordRecovery)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	router.Use(cors.Default())
	log.Fatal(router.Run())
}
