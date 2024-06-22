package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"uttc-backend/config"
	"uttc-backend/controller/like"
	"uttc-backend/controller/reply"
	"uttc-backend/controller/tweet"
	"uttc-backend/controller/user"

	"github.com/gin-gonic/gin"
)

func main() {

	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatalf("Error loading .env file: %v", err)
	//
	//}
	// DB接続のための準備
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	var dsn string
	dsn = fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)

	log.Println(dsn)
	// データベース接続の初期化
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Could not establish database connection: %v", err)
	}

	config.SetDB(db)

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
