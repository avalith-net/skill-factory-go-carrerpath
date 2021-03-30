package controllers

import (
	"net/http"
	"strconv"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/gin-gonic/gin"
)

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
