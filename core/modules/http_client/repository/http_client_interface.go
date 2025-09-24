package repository

import (
	"context"
	"net/http"

	dto2 "github.com/cde/go-example/core/modules/http_client/dto"
)

//go:generate mockgen -destination=mocks/mock_HttpClientRepository.go -package=mocks . HttpClientRepository
type (
	HttpClientRepository interface {
		EnableDebug() HttpClientRepository
		DisableDebug() HttpClientRepository
		Do(ctx context.Context, request *http.Request, headers map[string]string) (*dto2.ResponseByte, error)
	}
)