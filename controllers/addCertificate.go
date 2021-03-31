package controllers

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/jwt"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/gin-gonic/gin"
)

//CreatePath godoc
// @Summary Add a new certificate
// @Description allows you add a certificate to your skill
// @User get-struct-by-json
// @Param body-json body string true "Complete the json to create new path"
// @Param pathid query string true "pathID"
// @Param Authorization header string true "Token"
// @Accept json
// @Success 200 {string} string "Certificate added. Congratulations!"
// @Failure 400 {string} string "Something went wrong with the given data: "
// @Failure 400 {string} string "User path not related with user: "
// @Failure 400 {string} string "could not modify user's path"
// @Failure 400 {string} string "Couldnt get user by ID"
// @Failure 400 {string} string "could not notify the admin"
// @Failure default {string} string "error"
// @Router /addCertificate [patch]
func AddCertificate(c *gin.Context) {
	var userPath models.RelatadPath
	var notification models.Notification
	pathID := c.Query("pathid")

	if err := c.ShouldBind(&userPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Something went wrong with the given data: ": err.Error()})
		return
	}
	_, _, relationID, err := database.ConsultUserPath(jwt.UserID, pathID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"User path not related with user: ": err.Error()})
		return
	}
	if _, err := database.ModifyUserPath(userPath, relationID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"could not modify user's path ": err.Error()})
		return
	}

	for _, techSkill := range userPath.Path.TechnicalSkills {
		if techSkill.Blocked && techSkill.Certified {
			notification.TechSkill = append(notification.TechSkill, techSkill.Name)
		}
	}
	for _, softSkill := range userPath.Path.SoftSkills {
		if softSkill.Blocked && softSkill.Certified {
			notification.SoftSkill = append(notification.SoftSkill, softSkill.Name)
		}
	}
	user, err := database.GetUserById(jwt.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Couldnt get user by ID ": err.Error()})
		return
	}
	notification.Name = user.Name + " " + user.Surname
	notification.UserId = jwt.UserID
	notification.Message = "I want to validate this skill, please check my certification. Thanks!"
	if err := database.SendNotification(notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"could not notify the admin ": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Certificate added. Congratulations!")
}
