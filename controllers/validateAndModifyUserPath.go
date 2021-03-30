package controllers

import (
	"github.com/avalith-net/skill-factory-go-carrerpath/utils"
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

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
