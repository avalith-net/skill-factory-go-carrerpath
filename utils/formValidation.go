package utils

import (
	"fmt"
	"regexp"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
)

var (
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

//FormValidation is used to validate the register form
func FormValidation(user models.User) error {
	ErrorString = ""
	if len(user.Name) == 0 {
		AppendError("name is needed")
	}
	if len(user.Surname) == 0 {
		AppendError("surname is needed")
	}
	if len(user.Email) == 0 {
		AppendError("email is needed")
	} else if !emailRegex.MatchString(user.Email) {
		AppendError("not a valid email")
	}
	if len(ErrorString) != 0 {
		return fmt.Errorf(ErrorString)
	}
	return nil
}
