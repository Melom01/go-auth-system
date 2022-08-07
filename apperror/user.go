package apperror

import (
	"fmt"
	"net/http"
)

func ErrUserAlreadyExist(message string) AppCustomError {
	var (
		code = "ERR_USER_ALREADY_EXIST"
		msg  = fmt.Sprintf("%s: an error was thrown during the user creation process. The error was: '%s'.", code, message)
	)

	return AppCustomError{
		message:    msg,
		statusCode: http.StatusConflict,
		parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}
