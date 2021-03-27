package controllers

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

func ValidateAndModifyUserPath(c *gin.Context) {
	//	var userPath models.Path
	var path models.Path
	PathID := c.Query("pathid")
	UserId := c.Query("userid")

	//looking the userpath
	pathUserDb, err := database.UserPath.SummaryUserPath(PathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"User path not related with user: ": err.Error()})
		return
	}

	_, err = database.ConsultUserPath(UserId,PathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"User path not related with user: ": err.Error()})
		return
	}

	if err := c.ShouldBind(&path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"something went wrong with the given data: ": err.Error()})
		return
	}

	if _, err := database.EditUserPath(path, pathUserDb, PathID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"could not modify user's path ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ("Path has been modified"))
}

