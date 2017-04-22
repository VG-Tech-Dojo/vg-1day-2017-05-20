package httputil

type APIResult interface{}

type APIError struct {
	Message string `json:"message"`
}

func newAPIError(err error) *APIError {
	return &APIError{
		Message: err.Error(),
	}
}

type ErrorResponse struct {
	Result *APIResult `json:"result"`
	Error  *APIError  `json:"error"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Error: newAPIError(err),
	}
}
