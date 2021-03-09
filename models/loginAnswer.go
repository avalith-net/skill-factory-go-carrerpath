package models

//LoginAnswer has the token that returns the login
type LoginAnswer struct {
	Token string `json:"token,omitempty"`
}
