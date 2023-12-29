package common

type Response struct {
	// 应用状态码，0 表示成功，非 0 表示失败
	Code int `json:"code" example:"0"`
	// 易于阅读的信息
	Message string `json:"message"`
	// 数据，当 Code 不为 0 时，Data 为 nil
}

func NewSucceedResponse(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}

func NewErrorResponse(code int, err error) *Response {
	return &Response{
		Code:    code,
		Message: err.Error(),
	}
}
