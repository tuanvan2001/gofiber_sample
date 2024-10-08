package Http

import "github.com/gofiber/fiber/v2"

type HttpError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail,omitempty"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func CreateHttpError(code int, message string, detail ...interface{}) *HttpError {
	var detailValue interface{}
	if len(detail) > 0 {
		detailValue = detail[0]
	}

	return &HttpError{
		Code:    code,
		Message: message,
		Detail:  detailValue,
	}
}

func CreateHttpErrorValidate(detail interface{}) *HttpError {
	return CreateHttpError(fiber.StatusBadRequest, "Xác thực body thất bại.", detail)
}
