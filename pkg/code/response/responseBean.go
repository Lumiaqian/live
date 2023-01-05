package response

type ResponseSuccessBean struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
type NullJson struct{}

func Success(data any) *ResponseSuccessBean {
	return &ResponseSuccessBean{200, "OK", data}
}

type ResponseErrorBean struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func Error(errCode uint32, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg}
}
