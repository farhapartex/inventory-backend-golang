package model

type RegistrationInput struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
	IsCustomer bool `json:"is_customer"`
	IsSuperAdmin bool `json:"is_super_admin"`
}

type LoginInput struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}