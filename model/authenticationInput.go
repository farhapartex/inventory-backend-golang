package model

type AuthenticationInput struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}