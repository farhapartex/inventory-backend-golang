package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goupp-backend/model"
	"net/http"
)

func Register(context *gin.Context){
	var input model.AuthenticationInput
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