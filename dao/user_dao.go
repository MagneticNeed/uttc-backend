package dao

import (
	"database/sql"

	"uttc-backend/model"
)

// InsertUser inserts a new user into the database
func InsertUser(db *sql.DB, user model.User) error {
	query := "INSERT INTO users (id, username, email, password) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, user.ID, user.Username, user.Email, user.Password)
	return err
}

// GetUserByID retrieves a user by their ID from the database
func GetUserByID(db *sql.DB, userID string) (*model.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE id = ?"
	row := db.QueryRow(query, userID)

	var user model.User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no user found with the given ID
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail retrieves a user by their email from the database
func GetUserByEmail(db *sql.DB, email string) (*model.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = ?"
	row := db.QueryRow(query, email)

	var user model.User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no user found with the given email
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates the user's information in the database
func UpdateUser(db *sql.DB, user model.User) error {
	query := "UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?"
	_, err := db.Exec(query, user.Username, user.Email, user.Password, user.ID)
	return err
}

// DeleteUser deletes a user from the database by their ID
func DeleteUser(db *sql.DB, userID string) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, userID)
	return err
}
