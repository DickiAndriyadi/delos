package response

type ErrorResponse struct {
	ErrCode int   `json:"errorCode"`
	Message error `json:"message"`
}
