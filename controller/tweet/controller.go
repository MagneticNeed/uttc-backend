package tweet

import (
	"net/http"
	"time"

	"uttc-backend/model"
	tweetUsecase "uttc-backend/usecase/tweet" // エイリアスを使用
	"uttc-backend/util"

	"github.com/gin-gonic/gin"
)

// PostTweetController handles the creation of a new tweet
func PostTweetController(c *gin.Context) {
	var tweet model.Tweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tweet.ID = util.NewULID()
	tweet.CreatedAt = time.Now()

	if err := tweetUsecase.PostTweetUsecase(tweet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tweet)
}

// DeleteTweetController handles the deletion of a tweet
func DeleteTweetController(c *gin.Context) {
	tweetID := c.Param("id")

	if err := tweetUsecase.DeleteTweetUsecase(tweetID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetTweetByIDController handles retrieving a tweet by its ID
func GetTweetByIDController(c *gin.Context) {
	tweetID := c.Param("id")

	tweet, err := tweetUsecase.GetTweetByIDUsecase(tweetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if tweet == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"})
		return
	}

	c.JSON(http.StatusOK, tweet)
}

// GetAllTweetsController handles retrieving all tweets
func GetAllTweetsController(c *gin.Context) {
	tweets, err := tweetUsecase.GetAllTweetsUsecase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tweets)
}
