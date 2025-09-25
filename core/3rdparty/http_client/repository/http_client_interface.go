package repository

import (
	"context"
	"net/http"

	dto2 "github.com/cde/go-example/core/3rdparty/http_client/dto"
)

//go:generate mockgen -destination=mocks/mock_HttpClientInterface.go -package=mocks . HttpClientInterface
type (
	HttpClientInterface interface {
		EnableDebug() HttpClientInterface
		DisableDebug() HttpClientInterface
		Do(ctx context.Context, request *http.Request, headers map[string]string) (*dto2.ResponseByte, error)
	}
)
