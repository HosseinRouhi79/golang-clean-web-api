package http

type HTTPResponse struct {
	Result           string             `json:"result"`
	Success          string             `json:"success"`
	ResultCode       string             `json:"resultCode"`
	ValidationErrors *[]ValidationError `json:"validationErrors"`
	Error            any                `json:"error"`
}
