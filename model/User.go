package model

import (
	"github.com/goupp-backend/config"
	"golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
    "html"
    "strings"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"size:255;not null;" json:"first_name"`
	LastName  string `gorm:"size:255:not null;" json:"last_name"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"-"`
	IsSuperAdmin bool `json:"is_super_admin"`
	IsCustomer bool `json:"is_customer"`
}

// User methods

func (user *User) Save() (*User, error) {
	err := config.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}


func (user *User) BeforeSave(*gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.FirstName = html.EscapeString(strings.TrimSpace(user.FirstName))
	user.LastName = html.EscapeString(strings.TrimSpace(user.LastName))
	return nil
}


func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}


// DB functions for User

func FindUserByUsername(username string) (User, error) {
	var user User
	err := config.Database.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}