package factory

import (
	repository2 "github.com/cde/go-example/src/modules/user/repository"
	"github.com/cde/go-example/src/modules/user/usecase"
	"gorm.io/gorm"
)

func ResolveUserRepository(db *gorm.DB) repository2.UserInterface {
	//return repository.NewUserDummy()
	return repository2.NewUserMySQL(db)
}

func ResolveUserUseCase(userRepository repository2.UserInterface) usecase.UserUseCaseInterface {
	return usecase.NewUserUseCase(userRepository)
}
