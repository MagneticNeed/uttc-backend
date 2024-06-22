package model

import "time"

type Like struct {
	ID        string    `json:"id"`
	PostID    string    `json:"post_id"`
	LikedBy   string    `json:"liked_by"`
	CreatedAt time.Time `json:"created_at"`
	PostType  string    `json:"post_type"`
}
