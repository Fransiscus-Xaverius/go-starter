package http_client

import (
	"context"
	"io"
	"net/http"
	"time"

	appContext "github.com/cde/go-example/core/context"
	"github.com/google/uuid"
)

//go:generate mockgen -destination=mocks/mock_HttpClientRepository.go -package=mocks . HttpClientRepository
type (
	Response[T any] struct {
		Content    T
		StatusCode int
		Duration   int64 // in millisecond
		Header     http.Header
	}

	ResponseByte struct {
		Status   int
		Content  []byte
		Header   http.Header
		Duration time.Duration
	}

	HttpClientRepository interface {
		EnableDebug() HttpClientRepository
		DisableDebug() HttpClientRepository
		Do(ctx context.Context, request *http.Request, headers map[string]string) (*ResponseByte, error)
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

func (v httpClientRepository) Do(ctx context.Context, request *http.Request, headers map[string]string) (*ResponseByte, error) {
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
		requestId = uuid.New().String()
	}

	if headers != nil && len(headers) > 0 {
		reqHeaders := make(http.Header)
		for headerKey, headerValue := range headers {
			reqHeaders[headerKey] = []string{headerValue}
		}
		request.Header = reqHeaders
	}

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
			WithField("duration", duration).
			Info("end request")
	}

	return &ResponseByte{
		Status:   response.StatusCode,
		Content:  content,
		Header:   response.Header,
		Duration: duration,
	}, nil
}
