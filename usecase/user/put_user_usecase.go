package usecase

import (
	"uttc-backend/config"
	"uttc-backend/dao"
	"uttc-backend/model"
)

// UpdateUserUsecase handles the business logic for updating a user's information
func UpdateUserUsecase(user model.User) error {
	// Get the database connection
	db := config.GetDB()

	// Call the DAO function to update the user
	err := dao.UpdateUser(db, user)
	if err != nil {
		return err
	}

	return nil
}
