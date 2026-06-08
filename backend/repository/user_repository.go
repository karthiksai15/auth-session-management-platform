package repository

import (
	"auth-system/backend/config"
	"auth-system/backend/models"
	"database/sql"
)

// CreateUser inserts a new user into the users table.
// It also reads back the auto-generated id, created_at, and updated_at.
func CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (name, email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`
	err := config.DB.QueryRow(query, user.Name, user.Email, user.Password, user.Role).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	return err
}

// FindByEmail looks up a user by their email address.
// Returns nil (no error) if the user does not exist.
func FindByEmail(email string) (*models.User, error) {
	user := &models.User{}

	query := `
		SELECT id, name, email, password, role, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	err := config.DB.QueryRow(query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil // User not found — this is not an error, just return nil
	}
	if err != nil {
		return nil, err // Something went wrong with the query
	}

	return user, nil
}
