package factory

import (
	userRepository "github.com/cde/go-example/src/modules/user/repository"
	"github.com/cde/go-example/src/modules/user/usecase"
	"gorm.io/gorm"
)

func ResolveUserRepository(db *gorm.DB) userRepository.UserInterface {
	//return userRepository.NewUserDummy()
	return userRepository.NewUserMySQL(db)
	//return userRepository.NewUserPostgreSQL()
}

func ResolveUserUseCase(userRepository userRepository.UserInterface) usecase.UserInterface {
	return usecase.NewUserUseCase(userRepository)
}
