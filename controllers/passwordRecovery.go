package controllers

import (
	"fmt"
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/avalith-net/skill-factory-go-carrerpath/utils"
	"github.com/gin-gonic/gin"
)

//PasswordRecovery godoc
// @Summary recovery the password if dont remember
// @Description send email at the person what forgot the password
// @User get-struct-by-json
// @Accept  json
// @Produce html
// @Success 200 {string} html
// @Header 200 {string} Token "jwtKey"
// @Failure 400,404 {object} gin.H "invalid email, please write your email or the given email is not registered"
// @Failure default {object} gin.H "error processing password recovery"
// @Router /passwordRecovery [post]
func PasswordRecovery(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"something went wrong with the given data: ": err.Error()})
		return
	}

	if len(user.Email) < 1 {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("please write your email"))
		return
	}

	//Antes de enviar el mail, chequea si es un mail ya registrado, sino devuelve el error
	_, finded, ID := database.CheckUser.CheckUserAlreadyExists(user.Email)
	if !finded {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("the given email is not registered"))
		return
	}

	msj, newPass, err := utils.PasswordRecovery(user.Name, user.Email)
	user.Password = newPass
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid ": err.Error()})
		return
	}
	_, err = database.PasswordUpdate(user, ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"couldn't update password ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, msj)

}
