package errmsg

type Error struct {
	Code       int    `json:"code"`
	HTTPStatus int    `json:"http_status"`
	Message    string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func New(code int, httpStatus int, message string) Error {
	return Error{
		Code:       code,
		HTTPStatus: httpStatus,
		Message:    message,
	}
}

var (
	ErrNotFound = New(999, 404, "not found")
)
