package infrastructure

import (
	"app/src/infrastructure/event"
	"app/src/infrastructure/user"
)

var NewUserRepository func() user.UserRepository = user.NewUserRepositoryMem

func UseUserRepositoryMem() {
	NewUserRepository = user.NewUserRepositoryMem
}

var NewEventRepository func() event.EventRepository = event.NewEventRepositoryMem

func UseEventRepositoryMem() {
	NewEventRepository = event.NewEventRepositoryMem
}