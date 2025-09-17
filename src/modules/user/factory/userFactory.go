package userFactory

import (
	"github.com/cde/go-example/src/modules/user/repository"
	user "github.com/cde/go-example/src/modules/user/usecase"
)

func ResolveUserRepository() repository.UserInterface {
	return repository.NewUserDummy()
}

func ResolveUserUseCase() user.UserUseCaseInterface {
	userRepository := ResolveUserRepository()
	return user.NewUserUseCase(userRepository)
}
