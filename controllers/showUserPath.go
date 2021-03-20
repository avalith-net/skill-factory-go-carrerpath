package controllers

import (
	"fmt"
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/gin-gonic/gin"
)

//ShowUserPath
// @Summary shows a given user career path
// @Description get user careerpath
// @User get-struct-by-json
// @Accept  json
// @Produce json
// @Success 200 {string} http.StatusOK "success show careerpath"
// @Header 200 {string} Token "jwtKey"
// @Failure 400 {string} http.StatusBadRequest "missing id parameter"
// @Failure 404 {string} http.StatusNotFound "error occurred looking the path"
// @Failure default {string} http.StatusBadRequest "Error show user path "
// @Router /login [post]
func ShowUserPath(c *gin.Context) {
	ID := c.Query("id")
	if len(ID) < 1 {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("missing id parameter"))
		return
	}
	path, err := database.UserPath.SummaryUserPath(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error occurred looking the path": err.Error()})
		return
	}
	c.JSON(http.StatusOK, path)
}
