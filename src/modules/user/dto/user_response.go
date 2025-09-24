package dto

import (
	"github.com/cde/go-example/src/modules/user/entity"
)

type UserResponse struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u UserResponse) FromUserEntity(userEntity *entity.User) UserResponse {
	return UserResponse{
		ID:    userEntity.ID,
		Name:  userEntity.Name,
		Email: userEntity.Email,
	}
}
