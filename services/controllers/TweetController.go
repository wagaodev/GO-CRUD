package controllers

import (
	entities "api/services/entities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweets []entities.Tweet  
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (t *tweetController) FindAll(ctx *gin.Context){
	if t.tweets == nil {
		ctx.JSON(http.StatusNotFound, gin.H {
			"Atenção": "O usuário ainda não escreveu nenhum Tweet, seja o primeiro",
		})
		return
	}	
	ctx.JSON(http.StatusOK, t.tweets)
}

func (t *tweetController) Create(ctx *gin.Context){
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		return
	}

	t.tweets = append(t.tweets, *tweet)

	ctx.JSON(http.StatusOK, t.tweets)
}

func (t *tweetController) Delete(ctx *gin.Context){
	id := ctx.Param("id")

	for index, tweet := range t.tweets {
		if tweet.ID == id {
			t.tweets = append(t.tweets[:index], t.tweets[index+1:]...)
			ctx.JSON(http.StatusOK, gin.H {
				"Success": fmt.Sprintf("Tweet with ID %s successfully deleted", id),
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H {
		"Error": "I'M SORRY !! Tweet not found",
	})
}