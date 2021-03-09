package utils

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	minPassLength = 6
)

var (
	hasUpperCase, hasLowerCase bool
	errorString                string
)

//ValidatePassword validate the password
func ValidatePassword(password string) error {
	//Initialize variables
	hasUpperCase, hasLowerCase = false, false
	errorString = ""
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpperCase = true
		case unicode.IsLower(char):
			hasLowerCase = true
		}
	}
	if !hasLowerCase {
		appendError("lowercase letter missing")
	}
	if !hasUpperCase {
		appendError("uppercase letter missing")
	}
	if len(password) < minPassLength {
		appendError(fmt.Sprintf("password length must be at least %d", minPassLength))
	}
	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}
	return nil
}
func appendError(err string) {
	if len(strings.TrimSpace(errorString)) != 0 {
		errorString += ", " + err
	} else {
		errorString = err
	}
}
