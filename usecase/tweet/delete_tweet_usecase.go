package usecase

import (
	"uttc-backend/config"
	"uttc-backend/dao"
)

// DeleteTweetUsecase handles the business logic for deleting a tweet
func DeleteTweetUsecase(tweetID string) error {
	// Get the database connection
	db := config.GetDB()

	// Call the DAO function to delete the tweet
	err := dao.DeleteTweet(db, tweetID)
	if err != nil {
		return err
	}

	return nil
}
