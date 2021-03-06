package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/avalith-net/skill-factory-go-carrerpath/utils"
	"github.com/gin-gonic/gin"
)

//Parameters that separated by spaces
//. param name,param type,data type,is mandatory?,comment attribute(optional)

//Register godoc
// @Summary creates the user register in the db
// @Description ask for email and password for register in the app
// @User get-struct-by-json
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param name formData string true "User credentials"
// @Param password formData string true "Password"
// @Param email formData string true "email"
// @Param profilePhoto formData file true "profilePhoto"
// @Param surname formData string true "surname"
// @Param role formData string false "role"
// @Success 200 {string} Token "user successfully registered"
// @Header 200 {string} Token "jwtKey"
// @Failure 400,404 {string} string "something went wrong with the given data, error or the given email is already in use"
// @Failure default {string} string "error in register"
// @Router /register [post]
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
		c.JSON(http.StatusBadRequest, ("the given email is already in use"))
		return
	}
	//Validate password. It needs at least, 1 upper case letter and 1 lower case
	if err := utils.ValidatePassword(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	//Save profile picture
	var extension = strings.Split(user.ProfilePhoto.Filename, ".")[1]
	if err := c.SaveUploadedFile(user.ProfilePhoto, "uploads/profile/"+strings.TrimSpace(user.Name)+"."+extension); err != nil {
		log.Fatal(err)
	}
	_, status, err := database.RegisterService.InsertRegister(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error trying to register user: ": err.Error()})
		return
	}
	if !status {
		c.JSON(http.StatusBadRequest, ("failed to insert user register"))
		return
	}
	//here everything went well
	c.JSON(http.StatusCreated, "user successfully registered")
}
