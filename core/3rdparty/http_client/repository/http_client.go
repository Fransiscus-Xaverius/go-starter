package repository

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	dto2 "github.com/cde/go-example/core/3rdparty/http_client/dto"
	appContext "github.com/cde/go-example/core/context"
	"github.com/cde/go-example/core/vars"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type (
	httpClient struct {
		client *http.Client
		debug  bool
	}
)

func NewHttpClientRepository(client *http.Client) HttpClientInterface {
	return &httpClient{client: client}
}

func (v httpClient) EnableDebug() HttpClientInterface {
	v.debug = true
	return v
}

func (v httpClient) DisableDebug() HttpClientInterface {
	v.debug = false
	return v
}

func (v httpClient) Do(ctx context.Context, request *http.Request, headers map[string]string) (*dto2.ResponseByte, error) {
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
		logger = logger.WithField(vars.LoggerFieldRequestId, requestId)
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
		logger.Errorf("httpClient.do got err %s", err.Error())
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

	return &dto2.ResponseByte{
		StatusCode: response.StatusCode,
		Content:    content,
		Header:     response.Header,
		Duration:   duration,
	}, nil
}
