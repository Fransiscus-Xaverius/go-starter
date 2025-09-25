package repository

import (
	"context"

	"github.com/cde/go-example/src/modules/user/entity"
)

type userPostgreSQL struct {
}

func NewUserPostgreSQL() UserInterface {
	return &userPostgreSQL{}
}

func (u userPostgreSQL) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userPostgreSQL) GetByID(ctx context.Context, id int32) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userPostgreSQL) List(ctx context.Context, limit int, offset int) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}
