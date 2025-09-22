package dto

type (
	AuthorizeResponse struct {
		Status string `json:"status"`
	}

	AuthorizeResponseData struct {
		UserId     int64  `json:"user_id"`
		BusinessId *int64 `json:"business_id,omitempty"`
		Type       string `json:"type"`
	}
)
