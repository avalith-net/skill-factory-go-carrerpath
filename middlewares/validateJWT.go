package middlewares

import (
	"fmt"
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/routers"
	"github.com/gin-gonic/gin"
)

//ValidateJWT is a middleware that ckeck if we have a valid token
func ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, _, _, err := routers.ProcessToken(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error in token: ": err.Error()})
			c.AbortWithStatus(401)
			return
		}
		fmt.Println("validateJWT middleware")
		c.Next()
	}
}
