package model

import "time"

type Reply struct {
	ID        string    `json:"id"`
	ParentID  string    `json:"parent_id"`
	PostedBy  string    `json:"posted_by"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	LikeCount int       `json:"like_count" `
}
