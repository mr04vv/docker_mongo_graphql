package models

import (
	"errors"
	"github.com/globalsign/mgo/bson"
)

type User struct {
	Id      	bson.ObjectId `json:"_id" bson:"_id"`
	UserName    string `json:"userName" bson:"userName"`
	Description string `json:"description" bson:"description"`
	ImageUrl    string `json:"imageUrl" bson:"imageUrl"`
	Email       string `json:"email" bson:"email"`
}

// constructor
func NewUser(
	userName string,
	description string,
	imageUrl string,
	email string) (*User, error) {

	// validate
	if userName == "" {
		return &User{"", "", "", "", ""}, errors.New("userName is empty")
	}
	if description == "" {
		return &User{"", "", "", "", ""}, errors.New("description is empty")
	}
	if imageUrl == "" {
		return &User{"", "", "", "", ""}, errors.New("imageUrl is empty")
	}
	if email == "" {
		return &User{"", "", "", "", ""}, errors.New("email is empty")
	}

	return &User{
		Id:      bson.NewObjectId(),
		UserName:    userName,
		Description: description,
		ImageUrl:    imageUrl,
		Email:       email}, nil
}

func (user *User) Equals(other User) bool {
	if user.Id == other.Id {
		return true
	}
	return false
}