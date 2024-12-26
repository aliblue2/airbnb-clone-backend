package middleware

import (
	"net/http"

	"airbnb.com/airbnb/utils"
	"github.com/gin-gonic/gin"
)

func AuthenticateUser(context *gin.Context) {
	token := context.Request.Header.Get("token")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"message": "missed token.!"})
		return
	}

	userId, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token.!"})
		return
	}

	context.Set("id", userId)

}
