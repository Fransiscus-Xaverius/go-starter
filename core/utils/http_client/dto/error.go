package dto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type (
	ResponseErr struct {
		Content    map[string]any `json:"content,omitempty"`
		StatusCode int            `json:"status_code,omitempty"`
		Duration   time.Duration  `json:"duration,omitempty"`
		Header     http.Header    `json:"header,omitempty"`
	}
)

func (r ResponseErr) Error() string {
	marshal, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("%+v", r.Content)
	}

	return string(marshal)
}
