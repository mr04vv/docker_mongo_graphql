package user

import (
	"errors"
	"app/src/models"
	"app/src/conf"
	"gopkg.in/mgo.v2/bson"
)

type UserRepositoryMem struct {
	users []*models.User
}

func NewUserRepositoryMem() UserRepository {
	return &UserRepositoryMem{[]*models.User{}}
}

// store event to repository
func (self *UserRepositoryMem) Store(user *models.User) UserRepository {
	conf.ConnectDB()
	conf.GetCollection("users").Insert(user)
	self.users = append(self.users, user)
	return self
}

func (self UserRepositoryMem) FindById(userId string) (*models.User, error) {
	conf.ConnectDB()
	query := conf.GetCollection("users").Find(bson.M{})
	query.All(&self.users)
	for _, val := range self.users {
		if val.Id.Hex() == userId {
			return val, nil
		}
	}

	return nil, errors.New("user not found")
}

func (self UserRepositoryMem) UserList() []*models.User {
	conf.ConnectDB()
	query := conf.GetCollection("users").Find(bson.M{})
	query.All(&self.users)
	return self.users
}