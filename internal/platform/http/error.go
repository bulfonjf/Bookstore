package http

import (
	"bookstore/internal/apperror"
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func HandleErrorAsJson(w http.ResponseWriter, r *http.Request, code int, message string, err error) {
	if err != nil {
		logError(r, err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&ErrorResponse{Error: message})

}

var codes = map[string]int{
	apperror.ECONFLICT:       http.StatusConflict,
	apperror.EINVALID:        http.StatusBadRequest,
	apperror.ENOTFOUND:       http.StatusNotFound,
	apperror.ENOTIMPLEMENTED: http.StatusNotImplemented,
	apperror.EUNAUTHORIZED:   http.StatusUnauthorized,
	apperror.EINTERNAL:       http.StatusInternalServerError,
}

func ErrorStatusCode(code string) int {
	if v, ok := codes[code]; ok {
		return v
	}
	return http.StatusInternalServerError
}

func logError(r *http.Request, err error) {
	log.Printf("[http] error: %s %s: %s", r.Method, r.URL.Path, err)
}
