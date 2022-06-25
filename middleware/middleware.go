package middleware

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"sentinel/apperror"
	"sentinel/logger"
)

func LoggingMiddleware(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}

func RecoveryPanicMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			rec := recover()
			if rec != nil {
				err, ok := rec.(apperror.AppCustomError)
				if ok {
					logger.LogMessageInRed("(PANIC) -> " + err.Error())
					sendCustomError(w, err.Status(), err.Params())
				} else {
					logger.LogMessageInRed("(PANIC) -> this panic error was not a custom Sentinel application error")
					err = apperror.ErrServerError("")
					sendCustomError(w, err.Status(), err.Params())
				}
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func sendCustomError(w http.ResponseWriter, code int, param map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(param)
	if err != nil {
		logger.LogMessageInRed("Error during the encoding of JSON panic response.")
	}
}
