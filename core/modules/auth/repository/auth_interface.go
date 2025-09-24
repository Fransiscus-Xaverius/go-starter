package repository

import (
	"context"

	"github.com/cde/go-example/core/modules/auth/dto"
)

//go:generate mockgen -destination=mocks/mock_AuthRepositoryInterface.go -package=mocks . AuthRepositoryInterface
type AuthRepositoryInterface interface {
	Authorize(ctx context.Context, token string) (*dto.AuthorizeResponse, error)
}
