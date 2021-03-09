package jwt

import (
	"time"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/dgrijalva/jwt-go"
)

//GenerateJWT generates the JWT for the authentication
func GenerateJWT(user models.User) (string, error) {
	privateKey := []byte("SkillFactoryGoAdvance_Avalith")

	playload := jwt.MapClaims{
		"email":   user.Email,
		"name":    user.Name,
		"surname": user.Surname,
		"_id":     user.ID.Hex(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, playload)
	tokenStr, err := token.SignedString(privateKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
