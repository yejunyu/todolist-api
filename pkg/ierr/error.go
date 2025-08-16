package ierr

// APIError 是我们自定义的错误类型，包含了业务所需的信息
type APIError struct {
	HTTPStatus int    // HTTP 状态码
	Code       int    // 业务错误码
	Msg        string // 错误信息
}

// Error 实现 error 接口
func (e *APIError) Error() string {
	return e.Msg
}

// New 创建一个新的 APIError
func New(httpStatus, code int, msg string) *APIError {
	return &APIError{
		HTTPStatus: httpStatus,
		Code:       code,
		Msg:        msg,
	}
}

// 定义一些常用的业务错误
var (
	ErrInvalidInput       = New(400, 10001, "Invalid input parameters")
	ErrUserNotFound       = New(404, 20001, "User not found")
	ErrUsernameExists     = New(409, 20002, "Username already exists")
	ErrInvalidCredentials = New(401, 20003, "Invalid credentials")
	// ErrSystem ... 你可以定义更多业务错误
	ErrSystem = New(500, 500, "system error")
)
