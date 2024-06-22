package dao

import (
	"database/sql"
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
	query := "SELECT id, content, posted_by, created_at, like_count FROM tweets WHERE id = ?"
	row := db.QueryRow(query, tweetID)

	var tweet model.Tweet
	if err := row.Scan(&tweet.ID, &tweet.Content, &tweet.PostedBy, &tweet.CreatedAt, &tweet.LikeCount); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no tweet found with the given ID
		}
		return nil, err
	}
	return &tweet, nil
}

// GetAllTweets retrieves all tweets from the database
func GetAllTweets(db *sql.DB) ([]model.Tweet, error) {
	query := "SELECT id, content, posted_by, created_at, like_count FROM tweets"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tweets []model.Tweet
	for rows.Next() {
		var tweet model.Tweet
		if err := rows.Scan(&tweet.ID, &tweet.Content, &tweet.PostedBy, &tweet.CreatedAt, &tweet.LikeCount); err != nil {
			return nil, err
		}
		tweets = append(tweets, tweet)
	}
	return tweets, nil
}
