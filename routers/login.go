package routers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/jwt"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

//Login perform the user login
func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid email or password": err.Error()})
		return
	}
	if len(user.Email) == 0 {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("email is required"))
		return
	}
	userDB, exist := database.LoginTry(user.Email, user.Password)
	if !exist {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("invalid user or password"))
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
