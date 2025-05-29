package models

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTTokens struct {
	AccessToken string `json:"access_token"`
}