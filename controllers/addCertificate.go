package controllers

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/jwt"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

func AddCertificate(c *gin.Context) {
	var userPath models.RelatadPath
	var notification models.Notification
	pathID := c.Query("pathid")

	if err := c.ShouldBind(&userPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Something went wrong with the given data: ": err.Error()})
		return
	}
	_, _, relationID, err := database.ConsultUserPath(jwt.UserID, pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"User path not related with user: ": err.Error()})
		return
	}
	if _, err := database.ModifyUserPath(userPath, relationID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"could not modify user's path ": err.Error()})
		return
	}

	notification.UserId = jwt.UserID
	notification.Message = "I want to validate this skill, please check my certification. Thanks!"
	if  err := database.SendNotification(notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"could not notify the admin ": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Certificate added. Congratulations!")
}
