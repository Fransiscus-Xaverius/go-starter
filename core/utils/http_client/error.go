package http_client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	ResponseErr struct {
		Content    map[string]any
		StatusCode int
		Duration   int64 // in millisecond
		Header     http.Header
	}
)

func (r ResponseErr) Error() string {
	marshal, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("%+v", r.Content)
	}

	return string(marshal)
}
