package user

import (
	"errors"
	"app/src/models"
	"app/src/conf"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type UserRepositoryMem struct {
	users []*models.User
}

func NewUserRepositoryMem() UserRepository {
	conf.ConnectDB()
	var AllUsers []*models.User // フィールド定義時にこれを使う
	query := conf.GetCollection("users").Find(bson.M{})
	query.All(&AllUsers)
	fmt.Print(&AllUsers)
	return &UserRepositoryMem{AllUsers}
}

// store event to repository
func (self *UserRepositoryMem) Store(user *models.User) UserRepository {
	self.users = append(self.users, user)
	return self
}

func (self UserRepositoryMem) FindById(userId string) (*models.User, error) {
	for _, val := range self.users {
		if val.UserId == userId {
			return val, nil
		}
	}

	return nil, errors.New("user not found")
}

func (self UserRepositoryMem) UserList() []*models.User {
	return self.users
}