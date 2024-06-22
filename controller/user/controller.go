package user

import (
	"net/http"

	"uttc-backend/model"
	userUsecase "uttc-backend/usecase/user"

	"github.com/gin-gonic/gin"
)

// RegisterUserController handles the registration of a new user with Firebase authentication
func RegisterUserController(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Firebaseによる認証情報を含むユーザーIDを受け取る前提
	if user.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	if err := userUsecase.RegisterUserUsecase(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUserByIDController handles retrieving a user by its ID
func GetUserByIDController(c *gin.Context) {
	userID := c.Param("id")

	user, err := userUsecase.GetUserByIDUsecase(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsersController handles retrieving all users
func GetAllUsersController(c *gin.Context) {
	users, err := userUsecase.GetAllUsersUsecase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUserController handles updating a user's information
func UpdateUserController(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userUsecase.UpdateUserUsecase(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
