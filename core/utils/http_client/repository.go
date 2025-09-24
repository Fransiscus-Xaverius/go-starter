package http_client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	appContext "github.com/cde/go-example/core/context"
	"github.com/cde/go-example/core/middleware"
	"github.com/cde/go-example/core/utils/http_client/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

//go:generate mockgen -destination=mocks/mock_HttpClientRepository.go -package=mocks . HttpClientRepository
type (
	HttpClientRepository interface {
		EnableDebug() HttpClientRepository
		DisableDebug() HttpClientRepository
		Do(ctx context.Context, request *http.Request, headers map[string]string) (*dto.ResponseByte, error)
	}

	httpClientRepository struct {
		client *http.Client
		debug  bool
	}
)

func NewHttpClientRepository(client *http.Client) HttpClientRepository {
	return &httpClientRepository{client: client}
}

func (v httpClientRepository) EnableDebug() HttpClientRepository {
	v.debug = true
	return v
}

func (v httpClientRepository) DisableDebug() HttpClientRepository {
	v.debug = false
	return v
}

func (v httpClientRepository) Do(ctx context.Context, request *http.Request, headers map[string]string) (*dto.ResponseByte, error) {
	defer func() {
		if request != nil && request.Body != nil {
			_ = request.Body.Close()
		}
	}()

	var (
		contextBuilder = appContext.NewContextBuilder(ctx)
		logger         = contextBuilder.GetLogger()
	)

	requestId := contextBuilder.GetRequestId()
	if requestId == "" {
		requestId = utils.UUIDv4()
		logger = logger.WithField(middleware.LoggerRequestId, requestId)
	}

	reqHeaders := make(http.Header)
	reqHeaders[fiber.HeaderXRequestID] = []string{requestId}
	if headers != nil && len(headers) > 0 {
		for headerKey, headerValue := range headers {
			reqHeaders[headerKey] = []string{headerValue}
		}
		request.Header = reqHeaders
	}

	logger = logger.WithField("span_id", utils.UUIDv4())
	if request != nil && v.debug {
		logger.WithField("method", request.Method).
			WithField("url", request.URL.String()).
			WithField("headers", request.Header).
			Info("start request")
	}

	start := time.Now()
	response, err := v.client.Do(request)
	duration := time.Since(start)
	defer func() {
		if response != nil && response.Body != nil {
			_ = response.Body.Close()
		}
	}()

	if err != nil {
		logger.Errorf("httpClientRepository.do got err %s", err.Error())
		return nil, err
	}
	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response != nil && v.debug {
		logger.WithField("status_code", response.StatusCode).
			WithField("headers", response.Header).
			WithField("duration", fmt.Sprintf("%v", duration)).
			Info("end request")
	}

	return &dto.ResponseByte{
		StatusCode: response.StatusCode,
		Content:    content,
		Header:     response.Header,
		Duration:   duration,
	}, nil
}

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
	client HttpClientRepository,
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
