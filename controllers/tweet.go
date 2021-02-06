package controllers

import (
	"net/http"
	"strconv"

	models "app-v1/models"

	"github.com/gin-gonic/gin"
)

var myFirstTweet = models.Tweet{Id: 1, Title: "First tweet"}
var myTweets = []models.Tweet{myFirstTweet}

type TweetController struct{}

func (tweet TweetController) Home(context *gin.Context) {
	context.JSON(http.StatusNotFound, gin.H{"author": "Fatih Karadağ"})
}

func (tweet TweetController) Get(context *gin.Context) {
	context.JSON(http.StatusOK, myTweets)
}

func (tweet TweetController) GetById(context *gin.Context) {

	for _, myTweet := range myTweets {
		var id, err = strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "tweet id must be an integer"})
			return
		}

		if myTweet.Id == id {
			context.JSON(http.StatusOK, gin.H{
				"id":    myTweet.Id,
				"title":  myTweet.Title,
			})
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "tweet not found"})
	return
}

func (tweet TweetController) Post(context *gin.Context) {
	var json models.Tweet

	if err := context.ShouldBindJSON(&json); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	myTweets = append(myTweets, json)
	context.JSON(http.StatusCreated, gin.H{})
	return
}

func (tweet TweetController) Delete(context *gin.Context) {
	for index, myTweet := range myTweets {
		var id, err = strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "tweet id must be an integer"})
			return
		}

		if myTweet.Id == id {
			myTweets = append(myTweets[:index], myTweets[index+1:]...)
			context.JSON(http.StatusOK, gin.H{})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"error": "tweet not found"})
	return
}
