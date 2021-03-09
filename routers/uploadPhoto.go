package routers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

//UploadPhoto load the profile picture
func UploadPhoto(c *gin.Context) {
	//upload picture
	file, err := c.FormFile("profile")
	if err != nil {
		log.Fatal(err)
	}
	//extract the file extension
	var extension = strings.Split(file.Filename, ".")[1]
	var user models.User
	var status bool
	user.ProfilePhoto = UserID + "." + extension
	status, err = database.LoadProfilePicture(user, UserID)
	if err != nil || !status {
		c.JSON(http.StatusBadRequest, gin.H{"error while loading the image: ": err.Error()})
		return
	}
	err = c.SaveUploadedFile(file, "uploads/profile/"+UserID+"."+extension)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, fmt.Sprintf("file uploaded"))
}
