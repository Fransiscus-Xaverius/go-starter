package repository

import (
	"context"

	"github.com/cde/go-example/core/modules/auth/dto"
	"github.com/cde/go-example/core/utils/security"
)

type AyoAuthHttpClient struct {
	accessKey       *security.AccessKey
	headerAccessKey string
}

func (a AyoAuthHttpClient) Authorize(ctx context.Context, token string) (dto.AuthorizeResponse, error) {
	//TODO implement me
	panic("implement me")
}
