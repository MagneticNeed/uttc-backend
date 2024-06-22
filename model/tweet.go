package model

import "time"

type Tweet struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	PostedBy  string    `json:"posted_by"`
	CreatedAt time.Time `json:"created_at"`
	LikeCount int       `json:"like_count"`
}
