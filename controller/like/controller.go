package like

import (
	"net/http"
	"time"

	"database/sql"
	"uttc-backend/model"
	likeUsecase "uttc-backend/usecase/like" // エイリアスを使用
	"uttc-backend/util"

	"github.com/gin-gonic/gin"
)

// PostLikeController handles the creation of a new like
func PostLikeController(c *gin.Context) {
	var like model.Like
	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	like.ID = util.NewULID()
	like.CreatedAt = time.Now()

	if err := likeUsecase.PostLikeUsecase(like); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusConflict, gin.H{"error": "User already liked this post"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, like)
}

// DeleteLikeController handles the deletion of a like
func DeleteLikeController(c *gin.Context) {
	likeID := c.Param("id")
	var req struct {
		PostID string `json:"post_id"`
		UserID string `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := likeUsecase.DeleteLikeUsecase(likeID, req.PostID, req.UserID); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusConflict, gin.H{"error": "User has not liked this post"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like deleted successfully"})
}

// GetLikeController handles retrieving a like by its ID
func GetLikeController(c *gin.Context) {
	likeID := c.Param("id")

	like, err := likeUsecase.GetLikeByIDUsecase(likeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if like == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Like not found"})
		return
	}

	c.JSON(http.StatusOK, like)
}

// GetAllLikesController handles retrieving all likes
func GetAllLikesController(c *gin.Context) {
	likes, err := likeUsecase.GetLikesUsecase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, likes)
}
