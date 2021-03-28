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

	c.JSON(http.StatusCreated, "Certificate added. Congratulations!")
}
