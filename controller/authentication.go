package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goupp-backend/model"
	"github.com/goupp-backend/helper"
	"net/http"
)

func Register(context *gin.Context){
	var input model.RegistrationInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		FirstName: input.FirstName,
		LastName: input.LastName,
		Username: input.Username,
		Password: input.Password,
		IsCustomer: input.IsCustomer,
		IsSuperAdmin: input.IsSuperAdmin,
	}

	savedUser, err := user.Save()

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(context *gin.Context){
	var input model.LoginInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.FindUserByUsername(input.Username)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	err = user.ValidatePassword(input.Password)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
        return
    }

	token, err := helper.GenerateJWTToken(user)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"error": "Token generation failed, try again!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}