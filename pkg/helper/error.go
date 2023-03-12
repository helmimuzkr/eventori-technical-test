package helper

type ErrorCustom struct {
	Code    int
	Message string
}

func SetError(code int, message string) error {
	return ErrorCustom{
		Code:    code,
		Message: message,
	}
}

func (ce ErrorCustom) Error() string {
	return ce.Message
}

func ErrorResponse(err error) (int, map[string]string) {
	errorCustom := err.(ErrorCustom)
	return errorCustom.Code, map[string]string{"message": errorCustom.Message}
}
