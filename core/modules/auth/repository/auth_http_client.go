package repository

import (
	"context"

	"github.com/cde/go-example/core/modules/auth/dto"
	repository2 "github.com/cde/go-example/core/modules/http_client/repository"
	"github.com/cde/go-example/core/modules/security/repository"
)

type AyoAuthHttpClient struct {
	accessKey       *repository.AccessKey
	headerAccessKey string
	httpClient      repository2.HttpClientRepository
}

func (a AyoAuthHttpClient) Authorize(ctx context.Context, token string) (*dto.AuthorizeResponse, error) {
	panic("implement me")
}
