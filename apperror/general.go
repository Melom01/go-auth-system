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

func ErrBadSyntax() AppCustomError {
	var (
		code = "ERR_BAD_SYNTAX"
		msg  = fmt.Sprintf("%s: tried to make a request using body or form with missing or invalid fields.", code)
	)

	return AppCustomError{
		message:    msg,
		statusCode: http.StatusInternalServerError,
		parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}

func ErrMalformedBody() AppCustomError {
	var (
		code = "ERR_MALFORMED_BODY"
		msg  = fmt.Sprintf("%s: tried to make a request using a json body that could not be decoded.", code)
	)

	return AppCustomError{
		message:    msg,
		statusCode: http.StatusInternalServerError,
		parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}
