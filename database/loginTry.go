package database

import (
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"golang.org/x/crypto/bcrypt"
)

//LoginTry check the db login
func LoginTry(email, password string) (models.User, bool) {
	user, finded, _ := CheckUserAlreadyExists(email)
	if !finded {
		return user, finded
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	if err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes); err != nil {
		return user, false
	}
	return user, true

}
