package usecase

import (
	"uttc-backend/config"
	"uttc-backend/dao"
	"uttc-backend/model"
)

// PostTweetUsecase handles the business logic for creating a new tweet
func PostTweetUsecase(tweet model.Tweet) error {
	// Get the database connection
	db := config.GetDB()

	// Call the DAO function to insert the tweet
	err := dao.InsertTweet(db, tweet)
	if err != nil {
		return err
	}

	return nil
}
