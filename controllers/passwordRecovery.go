package controllers

import (
	"fmt"
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
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

	if len(user.Email) < 1 {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("please write your email"))
		return
	}

	//Antes de enviar el mail, chequea si es un mail ya registrado, sino devuelve el error
	_, finded, _ := database.CheckUserAlreadyExists(user.Email)
	if !finded {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("the given email is not registered"))
		return
	}

	msj, err := utils.PasswordRecovery(user.Name, user.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, msj)

}
