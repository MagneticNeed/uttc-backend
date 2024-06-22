package main

import (
	"log"
	"os"

	"uttc-backend/config"
	"uttc-backend/controller/like"
	"uttc-backend/controller/reply"
	"uttc-backend/controller/tweet"
	"uttc-backend/controller/user"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get environment variables for DB connection
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	// Initialize the database connection
	if err := config.InitDB(dbUser, dbPassword, dbName, dbHost); err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	r := gin.Default()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Routes
	r.POST("/likes", like.PostLikeController)
	r.GET("/likes", like.GetAllLikesController)
	r.GET("/likes/:id", like.GetLikeController)
	r.DELETE("/likes/:id", like.DeleteLikeController)

	r.POST("/replies", reply.PostReplyController)
	r.GET("/replies", reply.GetAllRepliesController)
	r.GET("/replies/:id", reply.GetReplyController)
	r.DELETE("/replies/:id", reply.DeleteReplyController)

	r.POST("/tweets", tweet.PostTweetController)
	r.DELETE("/tweets/:id", tweet.DeleteTweetController)
	r.GET("/tweets/:id", tweet.GetTweetByIDController)
	r.GET("/tweets", tweet.GetAllTweetsController)

	r.POST("/users", user.RegisterUserController)
	r.GET("/users", user.GetUserController)
	r.PUT("/users", user.UpdateUserController)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
