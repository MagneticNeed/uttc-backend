package usecase

import (
	"uttc-backend/config"
	"uttc-backend/dao"
	"uttc-backend/model"
)

// GetUserUsecase handles the business logic for retrieving a user by their ID
func GetUserByIDUsecase(userID string) (*model.User, error) {
	// Get the database connection
	db := config.GetDB()

	// Call the DAO function to retrieve the user
	user, err := dao.GetUserByID(db, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmailUsecase(userEmail string) (*model.User, error) {
	// Get the database connection
	db := config.GetDB()

	// Call the DAO function to retrieve the user
	user, err := dao.GetUserByEmail(db, userEmail)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetAllUsersUsecase retrieves all users
func GetAllUsersUsecase() ([]model.User, error) {
	db := config.GetDB()
	return dao.GetAllUsers(db)
}
