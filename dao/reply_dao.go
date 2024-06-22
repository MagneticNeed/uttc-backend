package dao

import (
	"database/sql"
	"time"
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
	query := "SELECT id, parent_id, posted_by, content, created_at, (SELECT COUNT(*) FROM likes WHERE post_id = replies.id) AS like_count FROM replies WHERE parent_id IS NOT NULL"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var replies []model.Reply
	for rows.Next() {
		var reply model.Reply
		var createdAtStr string
		var parentID sql.NullString
		if err := rows.Scan(&reply.ID, &parentID, &reply.PostedBy, &reply.Content, &createdAtStr, &reply.LikeCount); err != nil {
			return nil, err
		}
		if parentID.Valid {
			reply.ParentID = parentID.String
		} else {
			reply.ParentID = ""
		}
		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			return nil, err
		}
		reply.CreatedAt = createdAt
		replies = append(replies, reply)
	}
	return replies, nil
}

// GetReplyByID retrieves a reply by its ID from the database
func GetReplyByID(db *sql.DB, replyID string) (*model.Reply, error) {
	query := "SELECT id, parent_id, posted_by, content, created_at, (SELECT COUNT(*) FROM likes WHERE post_id = replies.id) AS like_count FROM replies WHERE id = ?"
	row := db.QueryRow(query, replyID)

	var reply model.Reply
	var createdAtStr string
	var parentID sql.NullString
	if err := row.Scan(&reply.ID, &parentID, &reply.PostedBy, &reply.Content, &createdAtStr, &reply.LikeCount); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no reply found with the given ID
		}
		return nil, err
	}
	if parentID.Valid {
		reply.ParentID = parentID.String
	} else {
		reply.ParentID = ""
	}
	createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		return nil, err
	}
	reply.CreatedAt = createdAt
	return &reply, nil
}

// GetRepliesByPostID retrieves replies for a given post
func GetRepliesByPostID(db *sql.DB, postID string) ([]model.Reply, error) {
	query := "SELECT id, parent_id, posted_by, content, created_at, (SELECT COUNT(*) FROM likes WHERE post_id = replies.id) AS like_count FROM replies WHERE parent_id = ?"
	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var replies []model.Reply
	for rows.Next() {
		var reply model.Reply
		var createdAtStr string
		var parentID sql.NullString
		if err := rows.Scan(&reply.ID, &parentID, &reply.PostedBy, &reply.Content, &createdAtStr, &reply.LikeCount); err != nil {
			return nil, err
		}
		if parentID.Valid {
			reply.ParentID = parentID.String
		} else {
			reply.ParentID = ""
		}
		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			return nil, err
		}
		reply.CreatedAt = createdAt
		replies = append(replies, reply)
	}
	return replies, nil
}
