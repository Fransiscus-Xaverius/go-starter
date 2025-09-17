package user

import (
	"context"
	"os/user"

	userDTO "github.com/cde/go-example/src/modules/user/dto"
	"github.com/cde/go-example/src/modules/user/entity"
	userRepository "github.com/cde/go-example/src/modules/user/repository"
)

type userUseCase struct {
	userRepository userRepository.UserInterface
}

func NewUserUseCase(userRepository userRepository.UserInterface) *userUseCase {
	return &userUseCase{userRepository: userRepository}
}

func (u userUseCase) CreateUser(ctx context.Context, request *userDTO.UserRequest) (*userDTO.UserResponse, error) {
	newRecord, err := u.userRepository.Create(
		ctx,
		&entity.User{
			Name:     request.Name,
			Email:    request.Email,
			Password: request.Password,
		},
	)
	if err != nil {
		return nil, err
	}

	return &userDTO.UserResponse{
		ID:    newRecord.ID,
		Name:  newRecord.Name,
		Email: newRecord.Email,
	}, err
}

func (u userUseCase) GetUser(ctx context.Context, id string) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) ListUsers(ctx context.Context, limit int, offset int) ([]user.User, error) {
	//TODO implement me
	panic("implement me")
}
