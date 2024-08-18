package Http

type HttpError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail,omitempty"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func CreateHttpError(code int, message string, detail interface{}) *HttpError {
	return &HttpError{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
}
