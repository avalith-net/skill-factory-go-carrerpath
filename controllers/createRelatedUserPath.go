package controllers

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

func CreateRelatedUser(c *gin.Context) {
	PathID := c.Query("pathid")
	UserId := c.Query("userid")
	if len(PathID) < 1 {
		c.JSON(http.StatusBadRequest, ("missing Path id"))
		return
	}
	if len(UserId) < 1 {
		c.JSON(http.StatusBadRequest, ("missing User id"))
		return
	}

	var rel models.RelatadPath
	rel.UserId = UserId
	rel.UserPathId = PathID

	path, err := database.UserPath.SummaryUserPath(PathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"No path in result": err.Error()})
		return
	}

	rel.Path = path

	status, err := database.CreateRelatedUserPath(rel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error in the created realtionship": err.Error()})
		return
	}
	if !status {
		c.JSON(http.StatusBadRequest, ("couldn´t create user path relation"))
		return
	}

	c.JSON(http.StatusCreated, ("User successfuly related to the selected path."))
}
