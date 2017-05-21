package httputil

// APIResult はexportする必要がないのであとで閉じます
type APIResult map[string]interface{}

// APIError も同様
type APIError struct {
	Message string `json:"message"`
}

func newAPIError(err error) *APIError {
	return &APIError{
		Message: err.Error(),
	}
}

// APIResponse は...このへんファイル分けるのがgoらしい
type APIResponse struct {
	Result *APIResult `json:"result"`
	Error  *APIError  `json:"error"`
}

// NewErrorResponse はエラーメッセージを含んだAPIResponse構造体のポインタを返します
func NewErrorResponse(err error) *APIResponse {
	return &APIResponse{
		Error: newAPIError(err),
	}
}
