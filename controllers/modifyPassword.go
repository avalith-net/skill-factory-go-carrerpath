package controllers

import (
	"fmt"
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/jwt"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/avalith-net/skill-factory-go-carrerpath/utils"
	"github.com/gin-gonic/gin"
)

//ModifyUserPassword godoc
// @Summary Modify user password
// @Description get the user and password update
// @User get-struct-by-json
// @Param Body body string true "Allows to complete the json for login"
// @Param Authorization header string true "Token"
// @Success 200 {string} Token "Password updated"
// @Failure 400,404 {string} string "something went wrong with the given data or couldn't update password"
// @Failure default {string} string "Invalid Password Update"
// @Router /modifyPassword [put]
func ModifyUserPassword(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"something went wrong with the given data: ": err.Error()})
		return
	}
	if err := utils.ValidatePassword(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	if _, err := database.PasswordUpdate(user, jwt.UserID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"couldn't update password ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Password updated"))
}
