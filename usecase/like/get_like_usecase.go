package usecase

import (
	"uttc-backend/config"
	"uttc-backend/dao"
	"uttc-backend/model"
)

// GetLikesUsecase handles the business logic for retrieving likes
func GetLikesUsecase() ([]model.Like, error) {
	// Get the database connection
	db := config.GetDB()

	// Call the DAO function to retrieve the likes
	likes, err := dao.GetAllLikes(db)
	if err != nil {
		return nil, err
	}

	return likes, nil
}

// GetLikeByIDUsecase retrieves a like by its ID
func GetLikeByIDUsecase(likeID string) (*model.Like, error) {
	db := config.GetDB()
	return dao.GetLikeByID(db, likeID)
}
