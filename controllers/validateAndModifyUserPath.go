package controllers

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/utils"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

//ValidateAndModifyUserPath godoc
// @Summary Validate and modify the user path
// @Description Validate and modify the user career path if you are the admin
// @User get-struct-by-json
// @Accept  json
// @Param body-json body string true "Allows to complete the json for login"
// @Param pathid query string true "pathID"
// @Param userid query string true "userID"
// @Param message query string true "message"
// @Param Authorization header string true "Token"
// @Success 200 {string} string "Path has been modified"
// @Failure 400 {string} string "Couldnt get user by ID "
// @Failure 400 {string} string "something went wrong with the given data: "
// @Failure 400 {string} string "User path not related with user: "
// @Failure 400 {string} string "could not modify user's path "
// @Failure 400 {string} string "could send feedback to user. "
// @Failure default {string} string "Error"
// @Router /validateOrModifyUserPath [put]
func ValidateAndModifyUserPath(c *gin.Context) {
	var userPath models.RelatadPath
	PathID := c.Query("pathid")
	UserId := c.Query("userid")
	message := c.Query("message")

	user, err := database.GetUserById(UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Couldnt get user by ID ": err.Error()})
		return
	}

	if err := c.ShouldBind(&userPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"something went wrong with the given data: ": err.Error()})
		return
	}

	_, _, relationID, err := database.ConsultUserPath(UserId, PathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"User path not related with user: ": err.Error()})
		return
	}

	if _, err := database.ModifyUserPath(userPath, relationID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"could not modify user's path ": err.Error()})
		return
	}

	if err := utils.SendFeedbackCertificated(user.Name, user.Email, message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"could send feedback to user. ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Path has been modified")
}
