package apperror

import (
	"fmt"
	"net/http"
)

func ErrGocloak(message string) AppCustomError {
	var (
		code = "ERR_GOCLOAK"
		msg  = fmt.Sprintf("%s: gocloak error. The error was: '%s'.", code, message)
	)

	return AppCustomError{
		message:    msg,
		statusCode: http.StatusInternalServerError,
		parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}
