package user

import (
	"app/src/models"
)

type UserRepository interface {
	Store(user *models.User) UserRepository
	FindById(userId string) (*models.User, error)
	UserList() []*models.User
}