package controllers

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/jwt"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

//AddSkillUserPath godoc
// @Summary add new skill to the user path
// @Description get user careerpath
// @User get-struct-by-json
// @Accept  json
// @Param body-json body string true "Allows to complete the json for login"
// @Param pathid query string true "Insert path id"
// @Param Authorization header string true "Token"
// @Success 201 {string} string "Skills added, the admin will check them, please wait..."
// @Failure 400 {string} string "missing path id parameter"
// @Failure 400 {string} string "Something went wrong with the given data: "
// @Failure 400 {string} string "You can´t change the description"
// @Failure 400 {string} string "To validate you need admin permission"
// @Failure 400 {string} string "Path not related with user:"
// @Failure 400 {string} string "Couldn´t add the skills to the path: "
// @Failure default {string} string "Error"
// @Router /addSkill [patch]
func AddSkillUserPath(c *gin.Context) {
	var path models.RelatadPath
	pathID := c.Query("pathid")
	if len(pathID) < 1 {
		c.JSON(http.StatusBadRequest, "missing path id parameter")
		return
	}
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

	_, err = database.AddSkillUserPath(path, originalPath, relationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Couldn´t add the skills to the path: ": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Skills added, the admin will check them, please wait...")
}
