package dao

import (
	"database/sql"
	"time"
	"uttc-backend/model"
)

// InsertTweet inserts a new tweet into the database
func InsertTweet(db *sql.DB, tweet model.Tweet) error {
	query := "INSERT INTO tweets (id, content, posted_by, created_at, like_count) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, tweet.ID, tweet.Content, tweet.PostedBy, tweet.CreatedAt, tweet.LikeCount)
	return err
}

// DeleteTweet deletes a tweet from the database
func DeleteTweet(db *sql.DB, tweetID string) error {
	query := "DELETE FROM tweets WHERE id = ?"
	_, err := db.Exec(query, tweetID)
	return err
}

// GetTweetByID retrieves a tweet by its ID from the database
func GetTweetByID(db *sql.DB, tweetID string) (*model.Tweet, error) {
	query := "SELECT id, posted_by, content, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at, like_count FROM tweets WHERE id = ?"
	row := db.QueryRow(query, tweetID)

	var tweet model.Tweet
	var createdAtStr string
	if err := row.Scan(&tweet.ID, &tweet.PostedBy, &tweet.Content, &createdAtStr, &tweet.LikeCount); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no tweet found with the given ID
		}
		return nil, err
	}

	// Parse the created_at string to time.Time
	createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		return nil, err
	}
	tweet.CreatedAt = createdAt

	return &tweet, nil
}

// GetAllTweets retrieves all tweets from the database
func GetAllTweets(db *sql.DB) ([]model.Tweet, error) {
	query := "SELECT id, posted_by, content, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at, like_count FROM tweets"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tweets []model.Tweet
	for rows.Next() {
		var tweet model.Tweet
		var createdAtStr string
		if err := rows.Scan(&tweet.ID, &tweet.PostedBy, &tweet.Content, &createdAtStr, &tweet.LikeCount); err != nil {
			return nil, err
		}

		// Parse the created_at string to time.Time
		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			return nil, err
		}
		tweet.CreatedAt = createdAt

		tweets = append(tweets, tweet)
	}
	return tweets, nil
}
