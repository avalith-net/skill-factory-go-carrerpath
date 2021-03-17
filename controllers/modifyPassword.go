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

	_, finded, ID := database.CheckUser.CheckUserAlreadyExists(jwt.Email)
	if !finded {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("the given email is not registered"))
		return
	}

	if _, err := database.PasswordUpdate(user, ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"couldn't update password ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Password updated"))
}
