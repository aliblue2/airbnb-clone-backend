package routes

import (
	"net/http"

	"airbnb.com/airbnb/models"
	"airbnb.com/airbnb/utils"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	user := models.User{}

	context.ShouldBindJSON(&user)

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "an error ocurred durig signup proccess"})
		return
	}
	user.Password = hashedPassword
	userId, err := user.Signup()

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"message": "email already taken"})
		return
	}

	user.Id = userId
	context.JSON(http.StatusCreated, gin.H{"message": "user successfully created.!"})
}

func Login(context *gin.Context) {
	user := models.UserLogin{}
	context.ShouldBindJSON(&user)
	retrivedPassword, err := user.ValidateUserCreadentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "cant find user with this creadentials"})
		return
	}

	valid := utils.ComparePassword(user.Password, retrivedPassword)

	if !valid {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "password is invalid"})
		return
	}

	token, err := utils.GenerateToken(user.Id, user.Email, retrivedPassword)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant create token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})

}
