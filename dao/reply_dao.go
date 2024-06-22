package dao

import (
	"database/sql"
	"uttc-backend/model"
)

// InsertReply inserts a new reply into the database
func InsertReply(db *sql.DB, reply model.Reply) error {
	query := "INSERT INTO replies (id, parent_id, posted_by, content, created_at) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, reply.ID, reply.ParentID, reply.PostedBy, reply.Content, reply.CreatedAt)
	return err
}

// DeleteReply deletes a reply from the database
func DeleteReply(db *sql.DB, replyID string) error {
	query := "DELETE FROM replies WHERE id = ?"
	_, err := db.Exec(query, replyID)
	return err
}

// GetAllReplies retrieves all replies from the database
func GetAllReplies(db *sql.DB) ([]model.Reply, error) {
	query := "SELECT id, parent_id, posted_by, content, created_at, (SELECT COUNT(*) FROM likes WHERE post_id = replies.id AND post_type = 'reply') AS like_count FROM replies WHERE parent_id IS NOT NULL"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var replies []model.Reply
	for rows.Next() {
		var reply model.Reply
		if err := rows.Scan(&reply.ID, &reply.ParentID, &reply.PostedBy, &reply.Content, &reply.CreatedAt, &reply.LikeCount); err != nil {
			return nil, err
		}
		replies = append(replies, reply)
	}
	return replies, nil
}

// GetRepliesByPostID retrieves replies for a given post
func GetRepliesByPostID(db *sql.DB, postID string) ([]model.Reply, error) {
	query := "SELECT id, parent_id, posted_by, content, created_at, (SELECT COUNT(*) FROM likes WHERE post_id = replies.id AND post_type = 'reply') AS like_count FROM replies WHERE parent_id = ?"
	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var replies []model.Reply
	for rows.Next() {
		var reply model.Reply
		if err := rows.Scan(&reply.ID, &reply.ParentID, &reply.PostedBy, &reply.Content, &reply.CreatedAt, &reply.LikeCount); err != nil {
			return nil, err
		}
		replies = append(replies, reply)
	}
	return replies, nil
}

// GetReplyByID retrieves a reply by its ID from the database
func GetReplyByID(db *sql.DB, replyID string) (*model.Reply, error) {
	query := "SELECT id, parent_id, posted_by, content, created_at, (SELECT COUNT(*) FROM likes WHERE post_id = replies.id) AS like_count FROM replies WHERE id = ?"
	row := db.QueryRow(query, replyID)

	var reply model.Reply
	if err := row.Scan(&reply.ID, &reply.ParentID, &reply.PostedBy, &reply.Content, &reply.CreatedAt, &reply.LikeCount); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no reply found with the given ID
		}
		return nil, err
	}
	return &reply, nil
}
