package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"size:255;not null;" json:"first_name"`
	LastName  string `gorm:"size:255:not null;" json:"last_name"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"-"`
	IsSuperAdmin bool `json:"is_super_admin"`
	IsCustomer bool `json:"is_customer"`
}