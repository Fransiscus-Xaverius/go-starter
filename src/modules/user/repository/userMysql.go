package repository

import (
	"context"

	entity "github.com/cde/go-example/src/modules/user/entity"
	"gorm.io/gorm"
)

type userMySQL struct {
	db *gorm.DB
}

func NewUserMySQL(db *gorm.DB) UserInterface {
	return &userMySQL{db: db}
}

func (u userMySQL) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userMySQL) GetByID(ctx context.Context, id int32) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userMySQL) List(ctx context.Context, limit int, offset int) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}
