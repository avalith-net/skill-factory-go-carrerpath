package controllers

import (
	"fmt"
	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/jwt"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddSkillUserPath (c *gin.Context){
	var path models.RelatadPath
	path.UserPathId =  c.Query("pathid")

	if err := c.ShouldBind(&path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Something went wrong with the given data: ": err.Error()})
		return
	}
	path.UserId = jwt.UserID

	_, err := database.ConsultUserPath(path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"User path not related with user: ": err.Error()})
		return
	}

	_, err = database.AddSkillUserPath(path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Couldn´t add the skills to the path: ": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, fmt.Sprintf("Skills added, the admin will check them, please wait..."))
}

