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
		err    error
	)

	logger.Info("userMySQL.Create")
	defer func() {
		if err != nil {
			logger.Errorf("userMySQL.Create: %v", err)
		}
	}()
	err = db.Create(user).Error
	return user, err
}

func (u userMySQL) GetByID(ctx context.Context, id int32) (*entity.User, error) {
	var (
		logger = appContext.LoggerFromContext(ctx)
		db     = u.db.WithContext(ctx)
		user   entity.User
		err    error
	)
	logger.Info("userMySQL.GetByID")
	defer func() {
		if err != nil {
			logger.Errorf("userMySQL.GetByID: %v", err)
		}
	}()

	err = db.Where("id=?", id).First(&user).Error
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
		err    error
	)

	logger.Info("userMySQL.List")
	defer func() {
		if err != nil {
			logger.Errorf("userMySQL.List: %v", err)
		}
	}()

	err = db.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}
