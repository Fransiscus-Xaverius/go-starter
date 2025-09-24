package http_client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func marshalToBuffer[T any](content T) (*bytes.Buffer, error) {
	marshal, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(marshal), nil
}

func unmarshalResponseToError(response *ResponseByte) error {
	if response == nil {
		return ResponseErr{
			Content: map[string]any{
				"message": errors.New("nil response"),
			},
		}
	}
	contentByte := response.Content

	responseErr := ResponseErr{
		Content:    nil,
		StatusCode: response.Status,
		Duration:   response.Duration.Milliseconds(),
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

func errToResponseError(err error, response *ResponseByte) *ResponseErr {
	responseErr := ResponseErr{
		Content: map[string]any{
			"message": err.Error(),
		},
	}
	if response == nil {
		return &responseErr
	}

	responseErr.Duration = response.Duration.Milliseconds()
	responseErr.StatusCode = response.Status
	responseErr.Header = response.Header
	return &responseErr
}

func send[ResT any](
	ctx context.Context,
	client HttpClientRepository,
	request *http.Request,
	headers map[string]string,
) (*Response[ResT], error) {
	resp, err := client.Do(ctx, request, headers)
	if err != nil {
		return nil, errToResponseError(err, resp)
	}
	var content ResT
	defaultReps := Response[ResT]{
		Content:    content,
		StatusCode: resp.Status,
		Duration:   resp.Duration.Milliseconds(),
		Header:     resp.Header,
	}

	if !(resp.Status >= 200 && resp.Status < 300) {
		return nil, unmarshalResponseToError(resp)
	}

	err = json.Unmarshal(resp.Content, &content)
	if err != nil {
		return nil, errToResponseError(err, resp)
	}

	defaultReps.Content = content
	return &defaultReps, nil
}
