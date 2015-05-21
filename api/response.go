package api

type ApiResponse struct {
	StatusCode int    `json:"status_code,omitempty"`
	Response   string `json:"response,omitempty"`
}
