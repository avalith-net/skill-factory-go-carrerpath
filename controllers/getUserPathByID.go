package controllers

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/gin-gonic/gin"
)

//GetUserPathByID godoc
// @Summary shows a complete user's career path
// @Description get user careerpath
// @User get-struct-by-json
// @Accept  json
// @Param pathid query string true "pathID"
// @Param userid query string true "userID"
// @Param Authorization header string true "Token"
// @Success 200 {string} string "status ok"
// @Failure 400 {string} string "missing path id parameter"
// @Failure 400 {string} string "missing user path id parameter"
// @Failure 400 {string} string "user not related with given path"
// @Failure default {string} string "Error"
// @Router /getUserPath [get]
func GetUserPathByID(c *gin.Context) {
	pathID := c.Query("pathid")
	userID := c.Query("userid")
	if len(pathID) < 1 {
		c.JSON(http.StatusBadRequest, "missing path id parameter")
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
