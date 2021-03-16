package jwt

import (
	"errors"
	"strings"

	"github.com/avalith-net/skill-factory-go-carrerpath/database"
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/dgrijalva/jwt-go"
)

//Email is a global variable
var Email string

//UserID is a global variable
var UserID string

//ProcessToken validates the token and returns true if it is
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	privateKey := []byte("SkillFactoryGoAdvance_Avalith")
	claims := &models.Claim{}
	splitToken := strings.Split(token, "bearer")
	//It has to return 2 elements
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Invalid token format")
	}
	token = strings.TrimSpace(splitToken[1])
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	//at this point, if we have a valid token, we need to check if the email coming with the token is valid
	if err == nil {
		_, finded, _ := database.CheckUser.CheckUserAlreadyExists(claims.Email)
		if finded == true {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, finded, UserID, nil
	}
	if !parsedToken.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}
