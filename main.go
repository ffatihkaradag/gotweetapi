package main

import (
	"github.com/gin-gonic/gin"

	controllers "app-v1/controllers"
)

func main() {
	app := gin.Default()
	var tweetController controllers.TweetController
	app.GET("/tweet", tweetController.Get)
	app.GET("/tweet/:id", tweetController.GetById)
	app.POST("/tweet", tweetController.Post)
	app.DELETE("/tweet/:id", tweetController.Delete)
	app.Run()
}
