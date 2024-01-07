package response

// Response is a struct to api response
type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse is a struct to api response when a error occurs
type ErrorResponse struct {
	Code         int64       `json:"code"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}
