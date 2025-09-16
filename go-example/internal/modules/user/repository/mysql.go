package userRepository

import (
	"encoding/json"
	"os"
	entity "go-example/internal/modules/user/entity"
)

type mysqlRepo struct{}

func NewMySQLRepo() Repository {
    return &mysqlRepo{}
}

func (r *mysqlRepo) Create(user entity.User) (entity.User, error) {
	var users []entity.User
	data, _ := os.ReadFile("users.json")
	json.Unmarshal(data, &users)
	users = append(users, user)
	data, _ = json.Marshal(users)
	os.WriteFile("users.json", data, 0644)
    return user, nil
}

func (r *mysqlRepo) GetByID(id string) (entity.User, error) {
    return entity.User{ID: id, Name: "John", Email: "john@example.com"}, nil
}

func (r *mysqlRepo) List() ([]entity.User, error) {
    return []entity.User{{ID: "1", Name: "John", Email: "john@example.com"}}, nil
}