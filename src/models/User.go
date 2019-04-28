package models

import (
	"errors"
	"github.com/google/uuid"
)

type User struct {
	UserId      string `json:"userId" bson:"userId"`
	UserName    string `json:"userName" bson:"userName"`
	Description string `json:"description" bson:"description"`
	ImageUrl    string `json:"imageUrl" bson:"imageUrl"`
	Email       string `json:"email" bson:"email"`
}

// constructor
func NewUser(
	userName string,
	description string,
	ImageUrl string,
	email string) (*User, error) {

	// validate
	if userName == "" {
		return &User{"", "", "", "", ""}, errors.New("userName is empty")
	}
	if description == "" {
		return &User{"", "", "", "", ""}, errors.New("description is empty")
	}
	if ImageUrl == "" {
		return &User{"", "", "", "", ""}, errors.New("ImageUrl is empty")
	}
	if email == "" {
		return &User{"", "", "", "", ""}, errors.New("email is empty")
	}

	return &User{
		UserId:      uuid.New().String(),
		UserName:    userName,
		Description: description,
		ImageUrl:    ImageUrl,
		Email:       email}, nil
}

func (user *User) Equals(other User) bool {
	if user.UserId == other.UserId {
		return true
	}
	return false
}