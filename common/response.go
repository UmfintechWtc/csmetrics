package common

type Response struct {
	// 应用状态码，0 表示成功，非 0 表示失败
	Code int `json:"Code" example:"0"`
	// Shell 返回码，0 表示成功，非 0 表示失败
	Cli int `json:"Cli" example:"255"`
	// 易于阅读的信息
	Message string `json:"Message"`
	// 数据，当 Code 不为 0 时，Data 为 nil
}

func NewSucceedResponse(code, cliCode int, message string) *Response {
	return &Response{
		Code:    code,
		Cli:     cliCode,
		Message: message,
	}
}

func NewErrorResponse(code, cliCode int, err error) *Response {
	return &Response{
		Code:    code,
		Cli:     cliCode,
		Message: err.Error(),
	}
}
