package model

type Response struct {
	Result string          `json:"result"`           // 结果值；0标识失败，1标识成功
	Msg    string          `json:"msg"`              // 响应信息
	Errors []ErrorResponse `json:"errors,omitempty"` // 错误信息
}

type ErrorResponse struct {
	ErrorCode     string `json:"error_code"`               // 错误代码
	ErrorMsg      string `json:"error_msg"`                // 错误信息
	ReferenceCode string `json:"reference_code,omitempty"` // 参考代码
}
