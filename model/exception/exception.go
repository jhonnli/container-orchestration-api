package exception

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (err Error) Error() string {
	return err.Message
}

func NewError(code, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}
