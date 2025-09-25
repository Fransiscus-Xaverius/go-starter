package repository

import (
	"context"

	repository2 "github.com/cde/go-example/core/3rdparty/http_client/repository"
	"github.com/cde/go-example/core/modules/auth/dto"
	"github.com/cde/go-example/core/security/repository"
)

type AyoAuthHttpClient struct {
	accessKey       *repository.AccessKeyInterface
	headerAccessKey string
	httpClient      repository2.HttpClientInterface
}

func (a AyoAuthHttpClient) Authorize(ctx context.Context, token string) (*dto.AuthorizeResponse, error) {
	panic("implement me")
}
