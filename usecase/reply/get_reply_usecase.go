package usecase

import (
	"uttc-backend/config"
	"uttc-backend/dao"
	"uttc-backend/model"
)

// GetAllRepliesUsecase handles the business logic for retrieving all replies
func GetAllRepliesUsecase() ([]model.Reply, error) {
	// Get the database connection
	db := config.GetDB()

	// Call the DAO function to retrieve the replies
	replies, err := dao.GetAllReplies(db)
	if err != nil {
		return nil, err
	}

	return replies, nil
}

// GetReplyByIDUsecase retrieves a reply by its ID
func GetReplyByIDUsecase(replyID string) (*model.Reply, error) {
	db := config.GetDB()
	return dao.GetReplyByID(db, replyID)
}
