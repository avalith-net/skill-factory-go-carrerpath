package models

//LoginAnswer has the token that login returns
type LoginAnswer struct {
	Token string `json:"token,omitempty"`
}
