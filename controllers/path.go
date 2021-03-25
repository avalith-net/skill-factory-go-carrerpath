package controllers

import (
	"fmt"
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

func Path(c *gin.Context) {
	var path models.Path

	if err := c.ShouldBind(&path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"something went wrong with the given data": err.Error()})
		return
	}
	if _, err := database.CreatePath(path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"could not create Path": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Path has been created"))
}
