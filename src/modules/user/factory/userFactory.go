package userFactory

import (
	"github.com/cde/go-example/src/modules/user/repository"
	user "github.com/cde/go-example/src/modules/user/usecase"
	"gorm.io/gorm"
)

func ResolveUserRepository(db *gorm.DB) repository.UserInterface {
	//return repository.NewUserDummy()
	return repository.NewUserMySQL(db)
}

func ResolveUserUseCase(userRepository repository.UserInterface) user.UserUseCaseInterface {
	return user.NewUserUseCase(userRepository)
}
