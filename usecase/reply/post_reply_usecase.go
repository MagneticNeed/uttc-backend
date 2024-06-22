package usecase

import (
	"uttc-backend/config"
	"uttc-backend/dao"
	"uttc-backend/model"
)

// PostReplyUsecase handles the business logic for creating a new reply
func PostReplyUsecase(reply model.Reply) error {
	// Get the database connection
	db := config.GetDB()

	// Call the DAO function to insert the reply
	err := dao.InsertReply(db, reply)
	if err != nil {
		return err
	}

	return nil
}
