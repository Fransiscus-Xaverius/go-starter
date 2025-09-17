package test

import (
	entity "github.com/cde/go-example/src/modules/user/entity"
	"testing"
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
