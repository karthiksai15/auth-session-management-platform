package services

import (
	"auth-system/backend/models"
	"auth-system/backend/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser validates the input, hashes the password, and saves the user to the DB.
// The role is always set to "USER" — no one can self-register as ADMIN.
func RegisterUser(name, email, password string) (*models.User, error) {
	// Step 1: Check if the email is already taken
	existingUser, err := repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Step 2: Hash the password using bcrypt
	// Cost 10 is the default — it's secure and not too slow
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	// Step 3: Build the user object — role is always "USER" on registration
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "USER",
	}

	// Step 4: Save the user to the database
	err = repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
