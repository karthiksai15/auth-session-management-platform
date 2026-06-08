package services

import (
	"auth-system/backend/config"
	"auth-system/backend/models"
	"auth-system/backend/repository"
	"auth-system/backend/utils"
	"context"
	"errors"
	"fmt"
	"time"

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

// LoginUser finds the user by email, checks the password, and returns JWT tokens.
func LoginUser(email, password string) (string, string, error) {
	// Step 1: Find the user by email
	user, err := repository.FindByEmail(email)
	if err != nil {
		return "", "", err
	}
	if user == nil {
		// Use a generic message — don't reveal whether email or password was wrong
		return "", "", errors.New("invalid email or password")
	}

	// Step 2: Compare the provided password with the stored bcrypt hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	// Step 3: Generate a short-lived access token (15 min) with userId and role
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	// Step 4: Generate a long-lived refresh token (7 days)
	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	// Step 5: Store the refresh token in Redis
	// Key format: refresh_token:{userId}  e.g. refresh_token:1
	// TTL: 7 days — same as the token expiry
	redisKey := fmt.Sprintf("refresh_token:%d", user.ID)
	err = config.RedisClient.Set(context.Background(), redisKey, refreshToken, 7*24*time.Hour).Err()
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
