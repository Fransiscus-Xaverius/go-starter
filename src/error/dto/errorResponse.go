package dto

type ErrorResponse struct {
	Message 	string `json:"message"`
	Details 	string `json:"details"`
	ErrorCode 	string `json:"error_code"`
}