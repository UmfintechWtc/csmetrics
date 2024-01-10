package common

type Response struct {
	// 应用状态码，0 表示成功，非 0 表示失败
	Code int `json:"Code" example:"0"`
	// 易于阅读的信息
	Message string `json:"Message"`
}

func NewSucceedResponse(code int, message string) *Response {
	return &Response{
		Code:    0,
		Message: message,
	}
}

func NewErrorResponse(code int, err error) *Response {
	return &Response{
		Code:    code,
		Message: err.Error(),
	}
}
