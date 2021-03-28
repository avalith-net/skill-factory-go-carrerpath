package app

import (
	"github.com/avalith-net/skill-factory-go-carrerpath/controllers"
	_ "github.com/avalith-net/skill-factory-go-carrerpath/docs"
	"github.com/avalith-net/skill-factory-go-carrerpath/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Career Path API
// @version 1.0
// @description This is an application that allows you to manage your career path
// @termsOfService http://swagger.io/terms/

// @contact.name Career Path API Support
// @contact.url http://www.swagger.io/support
// @contact.email careerpath.avalith@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func mapUrls() {
	router.Use(middlewares.CheckDB())
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/getPath", middlewares.ValidateJWT(), controllers.GetPathByID)
	router.GET("/getUserPath", middlewares.ValidateJWT(), controllers.GetUserPathByID)
	router.PUT("/passwordRecovery", controllers.PasswordRecovery)
	router.PUT("/modifyPassword", middlewares.ValidateJWT(), controllers.ModifyUserPassword)
	router.POST("/createPath", middlewares.CheckPermissions(), controllers.CreatePath)
	router.POST("/createRelatedUserPath", middlewares.CheckPermissions(), controllers.CreateRelatedUser)
	router.PATCH("/addCertificate", middlewares.ValidateJWT(), controllers.AddCertificate)
	router.PUT("/validateOrModifyUserPath", middlewares.CheckPermissions(), controllers.ValidateAndModifyUserPath)
	router.PATCH("/addSkill", middlewares.ValidateJWT(), controllers.AddSkillUserPath)

	//use ginSwagger middleware to serve the API docs
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	//http://localhost:8080/swagger/index.html

}
