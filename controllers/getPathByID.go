package controllers

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/gin-gonic/gin"
)

//GetPathByID
// @Summary shows a complete original career path
// @Description get user careerpath
// @User get-struct-by-json
// @Accept  json
// @Param id query string true "ID"
// @Param Authorization header string true "Token"
// @Success 200 {string} string "success show careerpath"
// @Header 200 {string} Token "jwtKey"
// @Failure 400 {string} string "missing id parameter"
// @Failure 404 {string} string "error occurred looking the path"
// @Failure default {string} string "Error show careerpath"
// @Router /getPath [get]
func GetPathByID(c *gin.Context) {
	ID := c.Query("id")
	if len(ID) < 1 {
		c.JSON(http.StatusBadRequest, ("missing id parameter"))
		return
	}
	path, err := database.Path.GetPathByID(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error occurred looking for the path": err.Error()})
		return
	}
	c.JSON(http.StatusOK, path)
}
