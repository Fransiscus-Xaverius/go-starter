package userRepository

import (
	entity "go-example/internal/modules/user/entity"
)

type mockRepo struct {
    users map[string]entity.User
}

func NewMockRepo() Repository {
    return &mockRepo{users: make(map[string]entity.User)}
}

func (r *mockRepo) Create(u entity.User) (entity.User, error) {
    r.users[u.ID] = u
    return u, nil
}

func (r *mockRepo) GetByID(id string) (entity.User, error) {
    return r.users[id], nil
}

func (r *mockRepo) List() ([]entity.User, error) {
    var list []entity.User
    for _, u := range r.users {
        list = append(list, u)
    }
    return list, nil
}