package middlewares

import (
	"fmt"
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/gin-gonic/gin"
)

//CheckDB middleware checking database connection
func CheckDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		if database.ConnectionCheck() == 0 {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("lost database connection"))
			c.AbortWithStatus(401)
			return
		}
		fmt.Println("checkDB middleware")
		c.Next()
		// next.ServeHTTP(c.Writer, c.Request)
	}
}
