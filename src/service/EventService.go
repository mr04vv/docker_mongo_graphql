package service

import (
	"app/src/models"
	"app/src/infrastructure"
)

func FindEventById(userId string) (*models.Event, error) {
	repository := infrastructure.NewEventRepository()
	return repository.FindById(userId)
}