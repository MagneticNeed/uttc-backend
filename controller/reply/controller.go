package reply

import (
	"net/http"
	"time"

	"uttc-backend/model"
	replyUsecase "uttc-backend/usecase/reply" // エイリアスを使用
	"uttc-backend/util"

	"github.com/gin-gonic/gin"
)

// PostReplyController handles the creation of a new reply
func PostReplyController(c *gin.Context) {
	var reply model.Reply
	if err := c.ShouldBindJSON(&reply); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reply.ID = util.NewULID()
	reply.CreatedAt = time.Now()

	if err := replyUsecase.PostReplyUsecase(reply); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, reply)
}

// DeleteReplyController handles the deletion of a reply
func DeleteReplyController(c *gin.Context) {
	replyID := c.Param("id")

	if err := replyUsecase.DeleteReplyUsecase(replyID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetReplyController handles retrieving a reply by its ID
func GetReplyController(c *gin.Context) {
	replyID := c.Param("id")

	reply, err := replyUsecase.GetReplyByIDUsecase(replyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if reply == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reply not found"})
		return
	}

	c.JSON(http.StatusOK, reply)
}

// GetAllRepliesController handles retrieving all replies
func GetAllRepliesController(c *gin.Context) {
	replies, err := replyUsecase.GetAllRepliesUsecase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, replies)
}
