package http_client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cde/go-example/core/modules/http_client/dto"
	"github.com/cde/go-example/core/modules/http_client/repository"
	"github.com/gofiber/fiber/v2"
)

func marshalToBuffer[T any](content T) (*bytes.Buffer, error) {
	marshal, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(marshal), nil
}

func unmarshalResponseToError(response *dto.ResponseByte) error {
	if response == nil {
		return dto.ResponseErr{
			Content: map[string]any{
				"message": errors.New("nil response"),
			},
		}
	}
	contentByte := response.Content

	responseErr := dto.ResponseErr{
		Content:    nil,
		StatusCode: response.StatusCode,
		Duration:   response.Duration,
		Header:     response.Header,
	}

	if contentByte == nil || len(contentByte) == 0 {
		return responseErr
	}

	errData := map[string]any{}
	err := json.Unmarshal(contentByte, &errData)
	if err != nil {
		responseErr.Content = map[string]any{
			"message": string(contentByte),
		}

		return responseErr
	}

	responseErr.Content = errData
	return responseErr
}

func errToResponseError(err error, response *dto.ResponseByte) *dto.ResponseErr {
	responseErr := dto.ResponseErr{
		Content: map[string]any{
			"message": err.Error(),
		},
	}
	if response == nil {
		return &responseErr
	}

	responseErr.Duration = response.Duration
	responseErr.StatusCode = response.StatusCode
	responseErr.Header = response.Header
	return &responseErr
}

func send[ResT any](
	ctx context.Context,
	client repository.HttpClientRepository,
	request *http.Request,
	headers map[string]string,
) (*dto.Response[ResT], error) {
	resp, err := client.Do(ctx, request, headers)
	if err != nil {
		return nil, errToResponseError(err, resp)
	}
	var content ResT
	if !(resp.StatusCode >= fiber.StatusOK && resp.StatusCode < fiber.StatusBadRequest) {
		return nil, unmarshalResponseToError(resp)
	}

	err = json.Unmarshal(resp.Content, &content)
	if err != nil {
		return nil, errToResponseError(err, resp)
	}

	return &dto.Response[ResT]{
		Content:    content,
		StatusCode: resp.StatusCode,
		Duration:   resp.Duration,
		Header:     resp.Header,
	}, nil
}

func Post[ReqT any, ResT any](ctx context.Context, client repository.HttpClientRepository, url string, payload ReqT, headers map[string]string) (*dto.Response[ResT], error) {
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

func Get[ResT any](ctx context.Context, client repository.HttpClientRepository, url string, queries map[string][]string, headers map[string]string) (*dto.Response[ResT], error) {
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

func Put[ReqT any, ResT any](ctx context.Context, client repository.HttpClientRepository, url string, payload ReqT, headers map[string]string) (*dto.Response[ResT], error) {
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

func Delete[ResT any](ctx context.Context, client repository.HttpClientRepository, url string, headers map[string]string) (*dto.Response[ResT], error) {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	return send[ResT](ctx, client, request, headers)
}
