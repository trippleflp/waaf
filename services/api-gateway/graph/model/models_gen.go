// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Token struct {
	Token string `json:"token"`
}

type UserLoginData struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password string  `json:"password"`
}

type UserRegistrationData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}