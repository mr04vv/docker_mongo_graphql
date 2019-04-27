package service

import (
	"app/src/models"
	"app/src/infrastructure"
)

func FindUserById(userId string) (*models.User, error) {
	repository := infrastructure.NewUserRepository()
	return repository.FindById(userId)
}