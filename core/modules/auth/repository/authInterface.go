package repository

import (
	"context"

	"github.com/cde/go-example/core/modules/auth/dto"
)

type AuthRepositoryInterface interface {
	Authorize(ctx context.Context, token string) (dto.AuthorizeResponse, error)
}
