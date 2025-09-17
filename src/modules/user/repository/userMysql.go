package repository

import (
	"context"
	"errors"

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
	err := u.db.Create(user).Error
	return user, err
}

func (u userMySQL) GetByID(ctx context.Context, id int32) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("id=?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, err
}

func (u userMySQL) List(ctx context.Context, limit int, offset int) ([]entity.User, error) {
	var users []entity.User
	err := u.db.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}
