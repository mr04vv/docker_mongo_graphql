package infrastructure

import (
	"app/src/infrastructure/user"
)

var NewUserRepository func() user.UserRepository = user.NewUserRepositoryMem

func UseUserRepositoryMem() {
	NewUserRepository = user.NewUserRepositoryMem
}