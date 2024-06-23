package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"uttc-backend/config"
	"uttc-backend/controller/like"
	"uttc-backend/controller/reply"
	"uttc-backend/controller/tweet"
	"uttc-backend/controller/user"

	"github.com/gin-contrib/cors"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {

	// // .envファイルの読み込み, デプロイ時はコメントアウト
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)

	// }
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

	// CORS設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://uttc-frontend-five.vercel.app/"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// プリフライトリクエストに対応するためのオプションルート
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(204)
	})

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// ルートの設定
	r.POST("/likes", func(c *gin.Context) {
		log.Println("POST /likes request received")
		like.PostLikeController(c)
	})
	r.GET("/likes", func(c *gin.Context) {
		log.Println("GET /likes request received")
		like.GetAllLikesController(c)
	})
	r.GET("/likes/:id", func(c *gin.Context) {
		log.Println("GET /likes/:id request received")
		like.GetLikeController(c)
	})
	r.DELETE("/likes/:id", func(c *gin.Context) {
		log.Println("DELETE /likes/:id request received")
		like.DeleteLikeController(c)
	})

	r.POST("/replies", func(c *gin.Context) {
		log.Println("POST /replies request received")
		reply.PostReplyController(c)
	})
	r.GET("/replies", func(c *gin.Context) {
		log.Println("GET /replies request received")
		reply.GetAllRepliesController(c)
	})
	r.GET("/replies/:id", func(c *gin.Context) {
		log.Println("GET /replies/:id request received")
		reply.GetReplyController(c)
	})
	r.DELETE("/replies/:id", func(c *gin.Context) {
		log.Println("DELETE /replies/:id request received")
		reply.DeleteReplyController(c)
	})

	r.POST("/tweets", func(c *gin.Context) {
		log.Println("POST /tweets request received")
		tweet.PostTweetController(c)
	})
	r.DELETE("/tweets/:id", func(c *gin.Context) {
		log.Println("DELETE /tweets/:id request received")
		tweet.DeleteTweetController(c)
	})
	r.GET("/tweets/:id", func(c *gin.Context) {
		log.Println("GET /tweets/:id request received")
		tweet.GetTweetByIDController(c)
	})
	r.GET("/tweets", func(c *gin.Context) {
		log.Println("GET /tweets request received")
		tweet.GetAllTweetsController(c)
	})

	r.POST("/users", func(c *gin.Context) {
		log.Println("POST /users request received")
		user.RegisterUserController(c)
	})
	r.GET("/users/:id", func(c *gin.Context) {
		log.Println("GET /users/:id request received")
		user.GetUserByIDController(c)
	})
	r.GET("/users", func(c *gin.Context) {
		log.Println("GET /users request received")
		user.GetAllUsersController(c)
	})
	r.PUT("/users", func(c *gin.Context) {
		log.Println("PUT /users request received")
		user.UpdateUserController(c)
	})
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
