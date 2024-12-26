package routes

import (
	"net/http"
	"strconv"

	"airbnb.com/airbnb/models"
	"github.com/gin-gonic/gin"
)

func CreateNewComment(context *gin.Context) {
	PlaceId, err := strconv.ParseInt(context.Param("placeId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cant add new comment"})
		return
	}

	userId := context.GetInt64("id")

	comment := models.Comment{}
	context.ShouldBindJSON(&comment)
	comment.User_id = userId
	comment.Place_id = PlaceId

	id, err := comment.AddNewComments()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant add new comment to this place"})
		return
	}

	comment.Id = id
	context.JSON(http.StatusCreated, gin.H{"message": "this comment successfully added"})
}

func GetAllCommentsByPlaceId(context *gin.Context) {
	PlaceId, err := strconv.ParseInt(context.Param("placeId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cant get all comments"})
		return
	}

	comments, err := models.GetCommentsPlace(PlaceId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant get all comments by place id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"comments": comments})

}
