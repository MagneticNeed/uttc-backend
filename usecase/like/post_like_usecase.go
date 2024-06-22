package usecase

import (
	"uttc-backend/config"
	"uttc-backend/dao"
	"uttc-backend/model"
)

// PostLikeUsecase handles the business logic for creating a new like
func PostLikeUsecase(like model.Like) error {
	// Get the database connection
	db := config.GetDB()

	// Call the DAO function to insert the like
	err := dao.InsertLike(db, like)
	if err != nil {
		return err
	}

	return nil
}
