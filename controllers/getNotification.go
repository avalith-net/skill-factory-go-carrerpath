package controllers

import (
	"net/http"
	"strconv"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/gin-gonic/gin"
)

//GetNotification godoc
// @Summary Get notifications
// @Description Get all the notifications from the users
// @User get-struct-by-json
// @Accept  json
// @Param page query string true "Insert the page"
// @Param Authorization header string true "Token"
// @Success 200 {string} string "status ok"
// @Failure 400 {string} string "Page must be a number"
// @Failure 400 {string} string "Coudn´t get the notifications"
// @Failure default {string} string "Error"
// @Router /notifications [get]
func GetNotification(c *gin.Context) {
	page := c.Query("page")
	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Page must be a number": err.Error()})
		return
	}
	pag := int64(pageTemp)
	notifications, err := database.GetNotifications(pag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Coudn´t get the notifications": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notifications)
}
