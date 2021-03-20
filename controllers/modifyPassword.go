package controllers

import (
	"fmt"
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/jwt"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

func ModifyUserPassword(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"something went wrong with the given data: ": err.Error()})
		return
	}

	if _, err := database.PasswordUpdate(user, jwt.UserID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"couldn't update password ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Password updated"))
}
