package repository

import (
	"context"
	"errors"

	appContext "github.com/cde/go-example/core/context"
	"github.com/cde/go-example/src/modules/user/entity"
	"gorm.io/gorm"
)

type userMySQL struct {
	db *gorm.DB
}

func NewUserMySQL(db *gorm.DB) UserInterface {
	return &userMySQL{db: db}
}

func (u userMySQL) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	var (
		logger = appContext.LoggerFromContext(ctx)
		db     = u.db.WithContext(ctx)
	)

	logger.Info("userMySQL.Create")
	err := db.Create(user).Error
	return user, err
}

func (u userMySQL) GetByID(ctx context.Context, id int32) (*entity.User, error) {
	var (
		logger = appContext.LoggerFromContext(ctx)
		db     = u.db.WithContext(ctx)
		user   entity.User
	)

	logger.Info("userMySQL.GetByID")
	err := db.Where("id=?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, err
}

func (u userMySQL) List(ctx context.Context, limit int, offset int) ([]entity.User, error) {
	var (
		logger = appContext.LoggerFromContext(ctx)
		db     = u.db.WithContext(ctx)
		users  []entity.User
	)

	logger.Info("userMySQL.List")
	err := db.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}
