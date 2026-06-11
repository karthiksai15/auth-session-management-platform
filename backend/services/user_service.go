package services

import (
	"auth-system/backend/models"
	"auth-system/backend/repository"
)

func GetProfile(userId int) (*models.User, error) {
	return repository.FindUserByID(userId)
}

func UpdateProfile(userId int, name string) error {
	return repository.UpdateUser(userId, name)
}

func GetAllUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}
