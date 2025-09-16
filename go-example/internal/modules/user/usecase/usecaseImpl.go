package user

import (

	"errors"
	userRepo "go-example/internal/modules/user/repository"
	userEntity "go-example/internal/modules/user/entity"

)
type useCase struct {
    repo userRepo.Repository
}

func NewUseCase(r userRepo.Repository) UserUseCase {
    return &useCase{repo: r}
}

func (uc *useCase) CreateUser(name, email string) (userEntity.User, error) {
    if name == "" || email == "" {
        return userEntity.User{}, errors.New("name and email are required")
    }
    newUser := userEntity.User{ Name: name, Email: email}
    newUser, err := uc.repo.Create(newUser)
    if err != nil {
        return userEntity.User{}, err
    }
    return newUser, nil
}

func (uc *useCase) GetUser(id string) (userEntity.User, error) {
    return uc.repo.GetByID(id)
}

func (uc *useCase) ListUsers() ([]userEntity.User, error) {
    return uc.repo.List()
}