package routes

import (
	"airbnb.com/airbnb/middleware"
	"github.com/gin-gonic/gin"
)

func RouterHandler(server *gin.Engine) {

	server.GET("/places", GetAllPlaces)
	server.GET("/places/:placeId", GetPlaceById)
	server.GET("/comments/:placeId", GetAllCommentsByPlaceId)
	server.GET("/features/:placeId", GetAllFeaturesByPlaceId)

	// auth required path
	authPath := server.Group("/")
	authPath.Use(middleware.AuthenticateUser)
	// places routes
	authPath.POST("/places", CreateNewPlace)
	authPath.PUT("/places/:placeId", EditPlaceById)
	authPath.DELETE("/places/:placeId", DeletePlaceById)

	//comment routes
	authPath.POST("/comments/:placeId", CreateNewComment)

	//place features routes
	authPath.POST("/features/:placeId", AddNewPlaceFeature)

	//user routes
	server.POST("/login", Login)
	server.POST("/signup", Signup)

	// owner routes
	server.POST("/owner/login", LoginOwner)
	server.POST("/owner/signup", SignupOwner)
}
