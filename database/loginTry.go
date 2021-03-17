package database

import (
	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	Login loginInterface = &login{}
)

type login struct{}

type loginInterface interface {
	LoginTry(string, string) (models.User, bool)
}

//LoginTry check the db login
func (login *login) LoginTry(email, password string) (models.User, bool) {
	user, finded, _ := CheckUser.CheckUserAlreadyExists(email)
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
