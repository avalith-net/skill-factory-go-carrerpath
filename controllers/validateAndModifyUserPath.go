package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//TO DO: crear nueva funcion para validar y modificar un path asociado a un ususario
func ValidateAndModifyUserPath(c *gin.Context) {
	//	var userPath models.Path
	// var path models.RelatadPath
	// PathID := c.Query("pathid")
	// UserId := c.Query("userid")

	//looking the userpath
	// originalPath, err := database.Path.GetPathByID(PathID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"User path not related with user: ": err.Error()})
	// 	return
	// }

	// _, relationID, err := database.ConsultUserPath(UserId, PathID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"User path not related with user: ": err.Error()})
	// 	return
	// }

	// if err := c.ShouldBind(&path); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"something went wrong with the given data: ": err.Error()})
	// 	return
	// }

	// if _, err := database.EditUserPath(path, originalPath, relationID); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"could not modify user's path ": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, ("Path has been modified"))
}
