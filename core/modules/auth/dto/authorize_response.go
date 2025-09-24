package dto

type (
	AuthorizeRequest struct {
		Token string `json:"token"`
	}

	AuthorizeResponse struct {
		Status string                `json:"status"`
		Data   AuthorizeResponseData `json:"data"`
	}

	AuthorizeResponseData struct {
		UserId     int64  `json:"user_id"`
		BusinessId *int64 `json:"business_id,omitempty"`
		Type       string `json:"type"`
	}
)
