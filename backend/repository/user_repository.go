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

// FindUserByID looks up a user by their ID.
// Used when refreshing a token — we need the current role from the DB.
func FindUserByID(id int) (*models.User, error) {
	user := &models.User{}

	query := `
		SELECT id, name, email, password, role, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	err := config.DB.QueryRow(query, id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetAllUsers returns all users from the database.
// Used only by the admin endpoint — normal users cannot call this.
func GetAllUsers() ([]models.User, error) {
	query := `
		SELECT id, name, email, role, created_at, updated_at
		FROM users
		ORDER BY id ASC
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Start with an empty slice — not nil — so JSON returns [] instead of null
	users := []models.User{}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
