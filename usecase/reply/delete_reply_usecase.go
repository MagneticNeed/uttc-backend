package usecase

import (
	"uttc-backend/config"
	"uttc-backend/dao"
)

// DeleteReplyUsecase handles the business logic for deleting a reply
func DeleteReplyUsecase(replyID string) error {
	// Get the database connection
	db := config.GetDB()

	// Call the DAO function to delete the reply
	err := dao.DeleteReply(db, replyID)
	if err != nil {
		return err
	}

	return nil
}
