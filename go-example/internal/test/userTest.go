package test

import (
	"testing"
	entity "go-example/internal/modules/user/entity"
)

func TestUserEntity(t *testing.T) {
	user := entity.User{
		ID:    "1",
		Name:  "John Doe",
		Email: "john@example.com",
	}

	if user.Name != "John Doe" {
		t.Errorf("Expected Name 'John Doe', got %s", user.Name)
	}
	if user.Email != "john@example.com" {
		t.Errorf("Expected Email 'john@example.com', got %s", user.Email)
	}
}