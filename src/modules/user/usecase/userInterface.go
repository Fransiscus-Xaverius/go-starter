package user

import (
	"context"
	"os/user"

	userDTO "github.com/cde/go-example/src/modules/user/dto"
)

type UserUseCaseInterface interface {
	CreateUser(ctx context.Context, request *userDTO.UserRequest) (*userDTO.UserResponse, error)
	GetUser(ctx context.Context, id string) (*user.User, error)
	ListUsers(ctx context.Context, limit int, offset int) ([]user.User, error)
}
