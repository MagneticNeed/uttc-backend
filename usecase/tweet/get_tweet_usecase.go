package usecase

import (
	"uttc-backend/config"
	"uttc-backend/dao"
	"uttc-backend/model"
)

// GetAllTweetsUsecase handles the business logic for retrieving all tweets
func GetAllTweetsUsecase() ([]model.Tweet, error) {
	// Get the database connection
	db := config.GetDB()

	// Call the DAO function to retrieve the tweets
	tweets, err := dao.GetAllTweets(db)
	if err != nil {
		return nil, err
	}

	return tweets, nil
}

// GetTweetByIDUsecase retrieves a tweet by its ID
func GetTweetByIDUsecase(tweetID string) (*model.Tweet, error) {
	db := config.GetDB()
	return dao.GetTweetByID(db, tweetID)
}
