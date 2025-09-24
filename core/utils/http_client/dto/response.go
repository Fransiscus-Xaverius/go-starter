package dto

import (
	"net/http"
	"time"
)

type (
	Response[T any] struct {
		Content    T             `json:"content,omitempty"`
		StatusCode int           `json:"status_code,omitempty"`
		Duration   time.Duration `json:"duration,omitempty"`
		Header     http.Header   `json:"header,omitempty"`
	}

	ResponseByte struct {
		StatusCode int           `json:"status_code,omitempty"`
		Content    []byte        `json:"content,omitempty"`
		Header     http.Header   `json:"header,omitempty"`
		Duration   time.Duration `json:"duration,omitempty"`
	}
)
