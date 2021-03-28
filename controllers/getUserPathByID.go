package controllers

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/gin-gonic/gin"
)

func GetUserPathByID(c *gin.Context) {
	pathID := c.Query("pathid")
	userID := c.Query("userid")
	if len(pathID) < 1 {
		c.JSON(http.StatusBadRequest, "missing user path id parameter")
		return
	}
	if len(userID) < 1 {
		c.JSON(http.StatusBadRequest, "missing user path id parameter")
		return
	}
	userPath, _, _, err := database.ConsultUserPath(userID, pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user not related with given path": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userPath)
}
