package routers

import (
	"fmt"
	"net/http"
	"regexp"
	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

//Register creates the user register in the db
func Register(c *gin.Context) {
	lenPassword := 6
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"something went wrong with the given data: ": err.Error()})
		return
	}
	//validate the emal
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(user.Email) == 0 {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("email is required"))
		return
	}

	if !emailRegex.MatchString(user.Email){
		c.JSON(http.StatusBadRequest, fmt.Sprintf("not a valid email"))
		return
	}
	//how many characters does the password need?
	if len(user.Password) < lenPassword {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("password needs at least 6 characters"))		
		return
	}
	//check if the user already exists
	_, finded, _ := database.CheckUserAlreadyExists(user.Email)
	if finded {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("the given email is already in use"))		
		return
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
