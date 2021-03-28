package controllers

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/jwt"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

func AddSkillUserPath(c *gin.Context) {
	var path models.RelatadPath
	pathID := c.Query("pathid")

	if err := c.ShouldBind(&path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Something went wrong with the given data: ": err.Error()})
		return
	}
	if len(path.Path.Description) > 0 {
		c.JSON(http.StatusBadRequest, "You can´t change the description")
		return
	}
	for _, block := range path.Path.TechnicalSkills {
		if !block.Blocked {
			c.JSON(http.StatusBadRequest, "To validate you need admin permission")
			return
		}
	}
	originalPath, _, relationID, err := database.ConsultUserPath(jwt.UserID, pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Path not related with user: ": err.Error()})
		return
	}

	// originalPath, err := database.Path.GetPathByID(pathID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"Couldn´t add the skills to the path: ": err.Error()})
	// 	return
	// }

	_, err = database.AddSkillUserPath(path, originalPath, relationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Couldn´t add the skills to the path: ": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Skills added, the admin will check them, please wait...")
}
