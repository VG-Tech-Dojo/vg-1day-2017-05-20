package httputil

type APIResult map[string]interface{}

type APIError struct {
	Message string `json:"message"`
}

func newAPIError(err error) *APIError {
	return &APIError{
		Message: err.Error(),
	}
}

type APIResponse struct {
	Result *APIResult `json:"result"`
	Error  *APIError  `json:"error"`
}

func NewErrorResponse(err error) *APIResponse {
	return &APIResponse{
		Error: newAPIError(err),
	}
}
