package errors

// 应用错误
type AppError struct {
	Msg      string `json:"msg"`
	Code     int    `json:"code"`
	HttpCode int    `json:"-"`
}

func (a AppError) Error() string {
	return a.Msg
}

func New(msg string) AppError {
	return AppError{
		Msg:      msg,
		Code:     FAIL,
		HttpCode: 500,
	}
}

// 设置错误消息
func (e *AppError) SetMsg(msg string) *AppError {
	e.Msg = msg
	return e
}

// 设置错误编码
func (e *AppError) SetCode(code int) *AppError {
	e.Code = code
	return e
}

// 设置http错误编码
func (e *AppError) SetHttpCode(httpCode int) *AppError {
	e.HttpCode = httpCode
	return e
}
