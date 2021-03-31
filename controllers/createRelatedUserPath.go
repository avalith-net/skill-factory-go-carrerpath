package controllers

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

//CreateRelatedUser godoc
// @Summary create the user's path relation
// @Description allows you to create a new user's path relation if you are an admin
// @User get-struct-by-json
// @Produce json
// @Success 200 {string} string "User successfuly related to the selected path"
// @Param pathid query string true "PathID"
// @Param userid query string true "UserID"
// @Param Authorization header string true "Token"
// @Header 200 {string} Token "jwtKey"
// @Failure 400 {string} string "missing User id"
// @Failure 400 {string} string "missing Path id"
// @Failure 400 {string} string "User is already related with the given path"
// @Failure 400 {string} string "Error in the created realtionship"
// @Failure 400 {string} string "couldn´t create user path relation"
// @Success 200 {string} string "User successfuly related to the selected path"
// @Failure default {string} string "error"
// @Router /createRelatedUserPath [post]
func CreateRelatedUser(c *gin.Context) {
	PathID := c.Query("pathid")
	UserID := c.Query("userid")
	if len(PathID) < 1 {
		c.JSON(http.StatusBadRequest, ("missing Path id"))
		return
	}
	if len(UserID) < 1 {
		c.JSON(http.StatusBadRequest, ("missing User id"))
		return
	}

	//consults if a relation between the user and the path already exist
	if _, status, _, _ := database.ConsultUserPath(UserID, PathID); status {
		c.JSON(http.StatusBadRequest, "User is already related with the given path")
		return
	}

	var rel models.RelatadPath
	rel.UserId = UserID
	rel.UserPathId = PathID

	path, err := database.Path.GetPathByID(PathID)
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

	c.JSON(http.StatusCreated, ("User successfuly related to the selected path"))
}
