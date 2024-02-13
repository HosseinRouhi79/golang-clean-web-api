package http

import "github.com/HosseinRouhi79/golang-clean-web-api/src/api/validation"

type HTTPResponse struct {
	Result           string                        `json:"result"`
	Success          string                        `json:"success"`
	ResultCode       string                        `json:"resultCode"`
	ValidationErrors *[]validation.ValidationError `json:"validationErrors"`
	Error            any                           `json:"error"`
}
