package apperror

import (
	"fmt"
	"net/http"
)

func ErrUnableToSaveUserData(message string) AppCustomError {
	var (
		code = "ERR_UNABLE_TO_SAVE_USER_DATA"
		msg  = fmt.Sprintf("%s: an error was thrown while saving user data. The error was: '%s'.", code, message)
	)

	return AppCustomError{
		message:    msg,
		statusCode: http.StatusInternalServerError,
		parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}

func ErrUnableToUpdateUserData(message string) AppCustomError {
	var (
		code = "ERR_UNABLE_TO_UPDATE_USER_DATA"
		msg  = fmt.Sprintf("%s: an error was thrown while updating user data. The error was: '%s'.", code, message)
	)

	return AppCustomError{
		message:    msg,
		statusCode: http.StatusInternalServerError,
		parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}

func ErrUnableToDeleteUserData(message string) AppCustomError {
	var (
		code = "ERR_UNABLE_TO_DELETE_USER_DATA"
		msg  = fmt.Sprintf("%s: an error was thrown while deleting user data. The error was: '%s'.", code, message)
	)

	return AppCustomError{
		message:    msg,
		statusCode: http.StatusInternalServerError,
		parameters: map[string]interface{}{
			"error_code": code,
		},
	}
}
