package repository

import (
	"context"

	"github.com/cde/go-example/src/modules/user/entity"
)

//go:generate mockgen -destination=mocks/mock_UserInterface.go -package=mocks . UserInterface
type UserInterface interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id int32) (*entity.User, error)
	List(ctx context.Context, limit int, offset int) ([]entity.User, error)
}
