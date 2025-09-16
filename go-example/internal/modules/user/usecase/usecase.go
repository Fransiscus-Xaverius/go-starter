package user

import "go-example/internal/modules/user/entity"

type UserUseCase interface {
    CreateUser(name, email string) (user.User, error)
    GetUser(id string) (user.User, error)
    ListUsers() ([]user.User, error)
}