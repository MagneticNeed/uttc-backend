package usecase

import (
	"uttc-backend/config"
	"uttc-backend/dao"
	"uttc-backend/model"
)

// RegisterUserUsecase handles the business logic for registering a new user
func RegisterUserUsecase(user model.User) error {
	// Get the database connection
	db := config.GetDB()

	// Call the DAO function to insert the user
	err := dao.InsertUser(db, user)
	if err != nil {
		return err
	}

	return nil
}
