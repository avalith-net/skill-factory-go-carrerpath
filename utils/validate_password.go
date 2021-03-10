package utils

import (
	"fmt"
	"unicode"
)

const (
	minPassLength = 6
)

var (
	hasUpperCase, hasLowerCase bool
)

//ValidatePassword validate the password
func ValidatePassword(password string) error {
	//Initialize variables
	hasUpperCase, hasLowerCase = false, false
	ErrorString = ""
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpperCase = true
		case unicode.IsLower(char):
			hasLowerCase = true
		}
	}
	if !hasLowerCase {
		AppendError("lowercase letter missing")
	}
	if !hasUpperCase {
		AppendError("uppercase letter missing")
	}
	if len(password) < minPassLength {
		AppendError(fmt.Sprintf("password length must be at least %d", minPassLength))
	}
	if len(ErrorString) != 0 {
		return fmt.Errorf(ErrorString)
	}
	return nil
}
