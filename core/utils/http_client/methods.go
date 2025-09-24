package http_client

import (
	"context"
	"net/http"

	"github.com/cde/go-example/core/utils/http_client/dto"
)

func Post[ReqT any, ResT any](ctx context.Context, client HttpClientRepository, url string, payload ReqT, headers map[string]string) (*dto.Response[ResT], error) {
	buffer, err := marshalToBuffer[ReqT](payload)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, buffer)
	if err != nil {
		return nil, err
	}

	return send[ResT](ctx, client, request, headers)
}

func Get[ResT any](ctx context.Context, client HttpClientRepository, url string, queries map[string][]string, headers map[string]string) (*dto.Response[ResT], error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if queries != nil {
		reqQueries := request.URL.Query()
		for queryKey, queryValue := range queries {
			for _, qv := range queryValue {
				reqQueries.Add(queryKey, qv)
			}
		}
		request.URL.RawQuery = reqQueries.Encode()
	}

	return send[ResT](ctx, client, request, headers)
}

func Put[ReqT any, ResT any](ctx context.Context, client HttpClientRepository, url string, payload ReqT, headers map[string]string) (*dto.Response[ResT], error) {
	buffer, err := marshalToBuffer[ReqT](payload)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPut, url, buffer)
	if err != nil {
		return nil, err
	}

	return send[ResT](ctx, client, request, headers)
}

func Delete[ResT any](ctx context.Context, client HttpClientRepository, url string, headers map[string]string) (*dto.Response[ResT], error) {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	return send[ResT](ctx, client, request, headers)
}
