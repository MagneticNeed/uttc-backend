package usecase

import (
	"uttc-backend/config"
	"uttc-backend/dao"
)

// DeleteLikeUsecase handles the business logic for deleting a like
func DeleteLikeUsecase(likeID string, postID string, userID string) error {
	// Get the database connection
	db := config.GetDB()

	// Call the DAO function to delete the like
	err := dao.DeleteLike(db, likeID, postID, userID)
	if err != nil {
		return err
	}

	return nil
}
