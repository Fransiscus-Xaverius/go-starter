package userRepository

import (
	entity "go-example/internal/modules/user/entity"
)

type Repository interface {
    Create(entity.User) (entity.User, error)
    GetByID(string) (entity.User, error)
    List() ([]entity.User, error)
}