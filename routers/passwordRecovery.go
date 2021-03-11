package routers

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/avalith-net/skill-factory-go-carrerpath/utils"
	"github.com/gin-gonic/gin"
)

func PasswordRecovery(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid email": err.Error()})
		return
	}

	msj, err := utils.PasswordRecovery(user.Name, user.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, msj)

}
