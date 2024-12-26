package routes

import (
	"net/http"
	"strconv"

	"airbnb.com/airbnb/models"
	"github.com/gin-gonic/gin"
)

func GetAllFeaturesByPlaceId(context *gin.Context) {

	placeId, err := strconv.ParseInt(context.Param("placeId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cant convet place id"})
		return
	}

	feature, err := models.GetPlaceFeatures(placeId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant get feature "})
		return
	}

	context.JSON(http.StatusOK, gin.H{"feature": feature})

}

func AddNewPlaceFeature(context *gin.Context) {
	placeId, err := strconv.ParseInt(context.Param("placeId"), 10, 64)
	ownerId := context.GetInt64("id")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant add new feature"})
		return
	}

	place, err := models.GetPlaceById(placeId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "we dont have place with this id"})
		return
	}

	if place.OwnerId != ownerId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "only creator of place can add feature"})
		return

	}

	featureItem := models.Feature{}
	context.ShouldBindJSON(&featureItem)
	featureItem.Place_id = placeId
	featureId, err := featureItem.AddNewPlaceFeature()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant add new feature to this place"})
		return
	}

	featureItem.Id = featureId
	context.JSON(http.StatusCreated, gin.H{"message": "feature successfully added"})

}
