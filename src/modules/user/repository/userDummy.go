package repository

import (
	"context"

	entity "github.com/cde/go-example/src/modules/user/entity"
)

type UserDummy struct{}

func NewUserDummy() UserInterface {
	return &UserDummy{}
}

func (u UserDummy) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	return &entity.User{
		ID:       1,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (u UserDummy) GetByID(ctx context.Context, id int32) (*entity.User, error) {
	return &entity.User{
		ID:       id,
		Name:     "",
		Email:    "",
		Password: "",
	}, nil
}

func (u UserDummy) List(ctx context.Context, limit int, offset int) ([]entity.User, error) {
	return []entity.User{
		{
			ID:       1,
			Name:     "",
			Email:    "",
			Password: "",
		},
	}, nil
}
