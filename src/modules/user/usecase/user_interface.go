package usecase

import (
	"context"

	"github.com/cde/go-example/src/modules/user/dto"
)

//go:generate mockgen -destination=mocks/mock_UserUseCaseInterface.go -package=mocks . UserUseCaseInterface
type UserUseCaseInterface interface {
	CreateUser(ctx context.Context, request *dto.UserRequest) (*dto.UserResponse, error)
	GetUser(ctx context.Context, id int32) (*dto.UserResponse, error)
	ListUsers(ctx context.Context, limit int, offset int) ([]dto.UserResponse, error)
}
