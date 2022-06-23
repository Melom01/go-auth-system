package apperror

import (
	"fmt"
	"net/http"
)

func ErrServerError(message string) AppCustomError {
	var (
		code = "ERR_INTERNAL_SERVER_ERROR"
		msg  = fmt.Sprintf("%s: an error was thrown during internal server processes. The error was: '%s'.", code, message)
	)

	return AppCustomError{
		message:    msg,
		statusCode: http.StatusInternalServerError,
		parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}
