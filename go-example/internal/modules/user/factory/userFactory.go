package userFactory

import (
	"go-example/internal/modules/user/repository"
)

func InitUserFactory() userRepository.Repository {
	return userRepository.NewMySQLRepo();
}