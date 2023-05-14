package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TweetController struct {
	Tweets []Tweet
}

type Tweet struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func NewTweetController() *TweetController {
	return &TweetController{}
}

func (tc *TweetController) FindAll(ctx *gin.Context) {
	if tc.Tweets == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Atenção": "O usuário ainda não escreveu nenhum Tweet, seja o primeiro",
		})
		return
	}
	ctx.JSON(http.StatusOK, tc.Tweets)
}

func (tc *TweetController) Create(ctx *gin.Context) {
	tweet := Tweet{}

	if err := ctx.BindJSON(&tweet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Falha ao fazer o bind do JSON",
		})
		return
	}

	tc.Tweets = append(tc.Tweets, tweet)

	ctx.JSON(http.StatusOK, tc.Tweets)
}

func (tc *TweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for index, tweet := range tc.Tweets {
		if tweet.ID == id {
			tc.Tweets = append(tc.Tweets[:index], tc.Tweets[index+1:]...)
			ctx.JSON(http.StatusOK, gin.H{
				"Success": fmt.Sprintf("Tweet with ID %s successfully deleted", id),
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"Error": "I'M SORRY !! Tweet not found",
	})
}

func TweetHandler(c *gin.Context) {
	tweetController := NewTweetController()

	switch c.Request.Method {
	case "GET":
		tweetController.FindAll(c)
	case "POST":
		tweetController.Create(c)
	case "DELETE":
		tweetController.Delete(c)
	default:
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"Error": "Método não permitido",
		})
	}
}
