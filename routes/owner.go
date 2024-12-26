package routes

import (
	"net/http"

	"airbnb.com/airbnb/models"
	"airbnb.com/airbnb/utils"
	"github.com/gin-gonic/gin"
)

func SignupOwner(context *gin.Context) {
	owner := models.Owner{}

	context.ShouldBindJSON(&owner)

	hashedPassword, err := utils.HashPassword(owner.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "an error ocurred durig signup proccess"})
		return
	}
	owner.Password = hashedPassword
	userId, err := owner.Signup()

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"message": "email already taken"})
		return
	}

	owner.Id = userId
	context.JSON(http.StatusCreated, gin.H{"message": "user successfully created.!"})
}

func LoginOwner(context *gin.Context) {
	owner := models.OwnerLogin{}
	context.ShouldBindJSON(&owner)
	retrivedPassword, err := owner.ValidateUserCreadentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "cant find user with this creadentials"})
		return
	}

	valid := utils.ComparePassword(owner.Password, retrivedPassword)

	if !valid {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "password is invalid"})
		return
	}

	token, err := utils.GenerateToken(owner.Id, owner.Phone, retrivedPassword)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant create token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})

}
