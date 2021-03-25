package middlewares

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/jwt"
	"github.com/avalith-net/skill-factory-go-carrerpath/utils"
	"github.com/gin-gonic/gin"
)

/*
Usage: You need to call CheckPermissions(string) in routes.go as this:

router.GET("/userpath", middlewares.CheckPermissions("ShowUserPath"), controllers.ShowUserPath)

This function 'l check if you are logged in and then 'l check if your assigned role has the
propper permissions to access the required endpoint.
The permissions are set in utils.Roles, if you need to add one please modify that file.
When you use CheckPermissions(string) you don't need to call middlewares.ValidateJWT()
because this function already do that.

The parameter CheckPermissions(string) is not case sensitive
*/

//CheckPermissions middleware
func CheckPermissions() gin.HandlerFunc {

	return func(c *gin.Context) {
		_, _, id, err := jwt.ProcessToken(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error in token: ": err.Error()})
			c.AbortWithStatus(401)
			return
		}
		user, err := database.GetUserById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error occurred looking for the user rol": err.Error()})
			c.AbortWithStatus(403)
			return
		}
		_, erro := utils.IsAllowed(user.Role)
		if erro != nil {
			c.JSON(http.StatusForbidden, gin.H{"Forbidden ": erro.Error()})
			c.AbortWithStatus(403)
			return
		}
		c.Next()
	}
}
