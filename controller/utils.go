package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"sentinel/apperror"
)

func SetJsonHeadersAndEncode(w http.ResponseWriter, v interface{}) {
	w.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		msg := fmt.Sprintln("Error during JSON encoding for response: ", err)
		apperror.ThrowError(apperror.ErrServerError(msg))
	}
}

func (ctrl *HTTPController) DecodeBody(r *http.Request, v interface{}) {
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		apperror.ThrowError(apperror.ErrMalformedBody())
	}

	val := validator.New()
	err = val.Struct(v)
	if err != nil {
		apperror.ThrowError(apperror.ErrBadSyntax())
	}
}
