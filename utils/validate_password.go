package utils

import (
	"fmt"
	"strings"
	"unicode"
)

//ValidatePassword validate the password
func ValidatePassword(password string) error {
	var hasUpperCase bool
	var hasLowerCase bool
	// var hasNumber bool
	// var hasSpecialChar bool
	const minPassLength = 6
	var errorString string

	for _, ch := range password {
		switch {
		// case unicode.IsNumber(ch):
		// 	hasNumber = true
		case unicode.IsUpper(ch):
			hasUpperCase = true
		case unicode.IsLower(ch):
			hasLowerCase = true
		// case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
		// 	hasSpecialChar = true
		case ch == ' ':
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}
	if !hasLowerCase {
		appendError("lowercase letter missing")
	}
	if !hasUpperCase {
		appendError("uppercase letter missing")
	}
	// if !hasNumber {
	// 	appendError("atleast one numeric character required")
	// }
	// if !hasSpecialChar {
	// 	appendError("special character missing")
	// }
	if !(minPassLength <= len(password)) {
		appendError(fmt.Sprintf("password length must be at least %d", minPassLength))
	}
	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}
	return nil
}
