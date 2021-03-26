package controllers

import (
	"fmt"
	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRelatedUser(c *gin.Context){
	PathID := c.Query("pathid")
	UserId := c.Query("userid")
	if len(PathID) < 1 {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("missing Path id"))
		return
	}
	if len(UserId) < 1 {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("missing User id"))
		return
	}

	var rel models.RelatadPath
	rel.UserId = UserId
	rel.UserPathId = PathID

	status, err := database.CreateRelatedUserPath(rel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error in the created realtionship": err.Error()})
		return
	}
	if status == false {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("couldn´t create user path relation"))
		return
	}

	c.JSON(http.StatusCreated,"Succesfully, Relation created")
}
