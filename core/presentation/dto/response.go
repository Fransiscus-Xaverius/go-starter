package dto

type (
	Response[T any] struct {
		Status    bool   `json:"status,omitempty"`
		Message   string `json:"message,omitempty"`
		ErrDetail string `json:"error_detail,omitempty"`
		ErrCode   string `json:"error_code,omitempty"`
		Data      T      `json:"data,omitempty"`
	}
)
