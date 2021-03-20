package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/avalith-net/skill-factory-go-carrerpath/utils"
	"github.com/gin-gonic/gin"
)

//Register godoc
// @Summary creates the user register in the db
// @Description ask for email and password for register in the app
// @User get-struct-by-json
// @Accept  json
// @Produce json
// @Success 200 {string} Token "user successfully registered"
// @Header 200 {string} Token "jwtKey"
// @Failure 400,404 {string} http.StatusBadRequest "something went wrong with the given data, error or the given email is already in use"
// @Failure default {string} http.StatusBadRequest "error in register"
// @Router /register [get]
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"something went wrong with the given data: ": err.Error()})
		return
	}
	if err := utils.FormValidation(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	//check if the user already exists
	_, finded, _ := database.CheckUser.CheckUserAlreadyExists(user.Email)
	if finded {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("the given email is already in use"))
		return
	}
	//Validate password. It needs at least, 1 upper case letter and 1 lower case
	if err := utils.ValidatePassword(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	//Save profile picture
	var extension = strings.Split(user.ProfilePhoto.Filename, ".")[1]
	if err := c.SaveUploadedFile(user.ProfilePhoto, "uploads/profile/"+user.Name+"."+extension); err != nil {
		log.Fatal(err)
	}
	_, status, err := database.RegisterService.InsertRegister(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error trying to register user: ": err.Error()})
		return
	}
	if !status {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to insert user register"))
		return
	}
	//here everything went well
	c.JSON(http.StatusCreated, "user successfully registered")
}
