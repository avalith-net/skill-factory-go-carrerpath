package controllers

import (
	"net/http"
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/jwt"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

//Login godoc
// @Summary Enter the system
// @Description get the email and password to access
// @User get-struct-by-json
// @id login
// @Param body-json body string true "Allows to complete the json for login"
// @Accept json
// @Success 200 {string} Token "Success Login"
// @Header 200 {string} Token "jwtKey"
// @Failure 400,404 {string} string "invalid login"
// @Failure default {string} string "error"
// @Router /login [post]
func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid email or password": err.Error()})
		return
	}
	if len(user.Email) == 0 {
		c.JSON(http.StatusBadRequest, ("email is required"))
		return
	}
	userDB, exist := database.Login.LoginTry(user.Email, user.Password)
	if !exist {
		c.JSON(http.StatusBadRequest, ("invalid user or password"))
		return
	}

	jwtKey, err := jwt.GenerateJWT(userDB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when trying to generate the token: ": err.Error()})
		return
	}
	generatedToken := models.LoginAnswer{
		Token: jwtKey,
	}

	c.JSON(http.StatusCreated, generatedToken)

	expirationTime := time.Now().Add(24 * time.Hour) //token expires in one day
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
