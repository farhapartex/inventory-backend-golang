package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsSuperAdmin bool `json:"is_super_admin"`
	IsCustomer bool `json:"is_customer"`
}