package controllers

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

func ModifyUserPath(c *gin.Context) {
	//	var userPath models.Path
	var path models.RelatadPath
	PathID := c.Query("pathid")
	UserId := c.Query("userid")

	path.UserId = UserId
	path.UserPathId = PathID
	//looking the userpath
	thepath, err := database.UserPath.SummaryUserPath(path.UserPathId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error occurred looking the path": err.Error()})
		return
	}

	if err := c.ShouldBind(&thepath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"something went wrong with the given data: ": err.Error()})
		return
	}

	_, err = database.ConsultUserPath(path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"User path not related with user: ": err.Error()})
		return
	}

	if _, err := database.EditUserPath(thepath, path.UserId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"could not modify user's path ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ("Path has been modified"))
}
