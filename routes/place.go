package routes

import (
	"net/http"
	"strconv"

	"airbnb.com/airbnb/models"
	"github.com/gin-gonic/gin"
)

func GetAllPlaces(context *gin.Context) {
	plaecs, err := models.GetAllPlaces()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cant get all places"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"places": &plaecs})

}

func GetPlaceById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("placeId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant get place with this id"})
		return
	}

	place, err := models.GetPlaceById(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cant get place with this id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"place": place})

}

func CreateNewPlace(context *gin.Context) {
	contextOwnerId := context.GetInt64("id")

	place := models.Place{}

	err := context.ShouldBindJSON(&place)
	place.OwnerId = contextOwnerId

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"message": "cant bind object"})
		return
	}

	placeId, err := place.CreateNewPlace()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant create new place"})
		return
	}

	place.Id = placeId
	context.JSON(http.StatusCreated, gin.H{"message": "place successfully created"})

}

func DeletePlaceById(context *gin.Context) {
	contextOwnerId := context.GetInt64("id")

	placeId, err := strconv.ParseInt(context.Param("placeId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "an error ocurred.!"})
		return
	}

	place, err := models.GetPlaceById(placeId)

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"message": "cant find place with this id"})
		return
	}

	if place.OwnerId != contextOwnerId {
		context.JSON(http.StatusConflict, gin.H{"message": "only creator can delete place"})
		return
	}

	err = models.DeletePlace(placeId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "an error ocurred during deleting place"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "place deleted successfully.!"})

}

func EditPlaceById(context *gin.Context) {
	contextOwnerId := context.GetInt64("id")

	placeId, err := strconv.ParseInt(context.Param("placeId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "an error ocurred.!"})
		return
	}

	place, err := models.GetPlaceById(placeId)

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"message": "cant find place with this id"})
		return
	}

	if place.OwnerId != contextOwnerId {
		context.JSON(http.StatusConflict, gin.H{"message": "only creator can update place"})
		return
	}

	bindingPlace := models.Place{}
	err = context.ShouldBindJSON(&bindingPlace)
	bindingPlace.OwnerId = contextOwnerId

	err = bindingPlace.EditPlaceById(placeId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant update place.!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": bindingPlace})

}
