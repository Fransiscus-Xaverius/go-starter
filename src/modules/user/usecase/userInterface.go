package usecase

import (
	"context"

	userDTO "github.com/cde/go-example/src/modules/user/dto"
)

//go:generate mockgen -destination=mocks/mock_UserUseCaseInterface.go -package=mocks . UserUseCaseInterface
type UserUseCaseInterface interface {
	CreateUser(ctx context.Context, request *userDTO.UserRequest) (*userDTO.UserResponse, error)
	GetUser(ctx context.Context, id int32) (*userDTO.UserResponse, error)
	ListUsers(ctx context.Context, limit int, offset int) ([]userDTO.UserResponse, error)
}
