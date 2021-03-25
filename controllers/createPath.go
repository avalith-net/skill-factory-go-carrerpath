package controllers

import (
	"fmt"
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

//CreatePath godoc
// @Summary Create new career path
// @Description allows you to create a new path if you are an admin
// @User get-struct-by-json
// @Param body-json body string true "Complete the json to create new path"
// @Param Authorization header string true "Token"
// @Accept json
// @Success 200 {string} string "Path has been created"
// @Failure 400 {string} string "could not create Path"
// @Failure default {string} string "error"
// @Router /createPath [post]
func CreatePath(c *gin.Context) {
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
