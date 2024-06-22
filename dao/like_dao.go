package dao

import (
	"database/sql"
	"uttc-backend/model"
)

// InsertLike inserts a new like into the database
func InsertLike(db *sql.DB, like model.Like) error {
	// Check if the user has already liked the post
	if hasLiked, _ := UserHasLiked(db, like.PostID, like.LikedBy); hasLiked {
		return sql.ErrNoRows // User already liked this post
	}

	// Insert the like into the likes table
	query := "INSERT INTO likes (id, post_id, liked_by, created_at, post_type) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, like.ID, like.PostID, like.LikedBy, like.CreatedAt, like.PostType)
	if err != nil {
		return err
	}

	// Update the like count in the tweets table
	updateQuery := "UPDATE tweets SET like_count = like_count + 1 WHERE id = ?"
	_, err = db.Exec(updateQuery, like.PostID)
	if err != nil {
		// If the tweet update fails, try to update the replies table
		updateQuery = "UPDATE replies SET like_count = like_count + 1 WHERE id = ?"
		_, err = db.Exec(updateQuery, like.PostID)
	}

	return err
}

// DeleteLike deletes a like from the database and updates the like count in the associated tweet or reply
func DeleteLike(db *sql.DB, likeID string, postID string, userID string) error {
	// Check if the user has not liked the post
	if hasLiked, _ := UserHasLiked(db, postID, userID); !hasLiked {
		return sql.ErrNoRows // User has not liked this post
	}

	// Delete the like from the likes table
	query := "DELETE FROM likes WHERE id = ?"
	_, err := db.Exec(query, likeID)
	if err != nil {
		return err
	}

	// Update the like count in the tweets table
	updateQuery := "UPDATE tweets SET like_count = like_count - 1 WHERE id = ?"
	_, err = db.Exec(updateQuery, postID)
	if err != nil {
		// If the tweet update fails, try to update the replies table
		updateQuery = "UPDATE replies SET like_count = like_count - 1 WHERE id = ?"
		_, err = db.Exec(updateQuery, postID)
	}
	return err
}

// UserHasLiked checks if a user has already liked a post
func UserHasLiked(db *sql.DB, postID string, userID string) (bool, error) {
	query := "SELECT COUNT(*) FROM likes WHERE post_id = ? AND liked_by = ?"
	var count int
	err := db.QueryRow(query, postID, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetAllLikes(db *sql.DB) ([]model.Like, error) {
	query := "SELECT id, post_id, liked_by, created_at, post_type FROM likes"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []model.Like
	for rows.Next() {
		var like model.Like
		if err := rows.Scan(&like.ID, &like.PostID, &like.LikedBy, &like.CreatedAt, &like.PostType); err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	return likes, nil
}

// GetLikeByID retrieves a like by its ID from the database
func GetLikeByID(db *sql.DB, likeID string) (*model.Like, error) {
	query := "SELECT id, post_id, liked_by, created_at, post_type FROM likes WHERE id = ?"
	row := db.QueryRow(query, likeID)

	var like model.Like
	if err := row.Scan(&like.ID, &like.PostID, &like.LikedBy, &like.CreatedAt, &like.PostType); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no like found with the given ID
		}
		return nil, err
	}
	return &like, nil
}

// CountLikesByPostID counts the number of likes for a given post
func CountLikesByPostID(db *sql.DB, postID string) (int, error) {
	query := "SELECT COUNT(*) FROM likes WHERE post_id = ?"
	var count int
	err := db.QueryRow(query, postID).Scan(&count)
	return count, err
}
