package app

import (
	"github.com/avalith-net/skill-factory-go-carrerpath/controllers"
	"github.com/avalith-net/skill-factory-go-carrerpath/middlewares"
)

func mapUrls() {
	router.Use(middlewares.CheckDB())
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/userpath", middlewares.ValidateJWT(), controllers.ShowUserPath)
	router.POST("/passwordRecovery", controllers.PasswordRecovery)
}
