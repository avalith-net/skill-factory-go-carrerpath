package utils

import "strings"

var (
	//ErrorString is used to append err messages
	ErrorString string
)

//AppendError appends the error messages
func AppendError(err string) {
	if len(strings.TrimSpace(ErrorString)) != 0 {
		ErrorString += ", " + err
	} else {
		ErrorString = err
	}
}
