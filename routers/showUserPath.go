package routers

import (
	"fmt"
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/gin-gonic/gin"
)

//ShowUserPath shows a given user career path
func ShowUserPath(c *gin.Context) {
	ID := c.Query("id")
	if len(ID) < 1 {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("missing id parameter"))
		return
	}
	path, err := database.SummaryUserPath(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error occurred looking the path": err.Error()})
		return
	}
	c.JSON(http.StatusOK, path)
}
