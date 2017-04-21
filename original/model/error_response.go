package model

type APIResponse struct {
	Result *APIResult `json:"result"`
	Error  *APIError  `json:"error"`
}

type APIResult struct {
	Message *Message `json:"message"`
}

type APIError struct {
	Message string `json:"message"`
}
