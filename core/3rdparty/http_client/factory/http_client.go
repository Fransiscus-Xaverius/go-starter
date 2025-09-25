package factory

import (
	"net/http"
	"time"

	"github.com/cde/go-example/config"
	"github.com/cde/go-example/core/3rdparty/http_client/repository"
	v1 "github.com/cde/go-example/core/3rdparty/http_client/repository/v1"
)

func MakeHttpClientV1(cfg *config.Config) *v1.HttpClient {
	client := v1.NewHttpClientRepository(&http.Client{
		Timeout: time.Duration(cfg.AyoAuthTimeout) * time.Second,
	})
	client.EnableDebug()
	return client
}

func ResolveHttpClientRepository(cfg *config.Config) repository.HttpClientInterface {
	return MakeHttpClientV1(cfg)
}
