package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/LochanRn/tiny-url-server/utils/logger"

	"net/http"

	"github.com/go-errors/errors"
)

const (
	AuthHeader = "Authorization"

	versionV1 = "/v1"
)

var WhiteListedApis = []string{
	versionV1 + "/system/authenticate",
	versionV1 + "/system/health/ping",
	versionV1 + "/system/config",
}

func GetStackErrorResponse(stackTrace string, v interface{}) error {
	var err error
	switch cause := v.(type) {
	case string:
		err = errors.New(cause)
	case error:
		err = cause
	default:
		err = errors.New("server panic")
	}
	return err
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				stackTrace := fmt.Sprintf("PANIC: %v\nStack Trace:\n%s\n", rec, errors.Wrap(rec, 0).ErrorStack())
				_, isHerr := rec.(error)
				if !isHerr {
					fmt.Print(stackTrace)
				}
				serverError := GetStackErrorResponse(stackTrace, rec)
				toJSONError(w, http.StatusInternalServerError, serverError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func toJSONError(w http.ResponseWriter, status int, e error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	type Err struct {
		Code  int
		Error string
	}
	er := Err{Code: status}
	if e != nil {
		er.Error = e.Error()
	}
	err := json.NewEncoder(w).Encode(er)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed with error while encoding json errors %v", err.Error()))
	}
}
