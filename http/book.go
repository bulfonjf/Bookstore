package http

import (
	"bookstore/application"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) registerBookRoutes(r *mux.Router) {
	r.HandleFunc("/books", s.createBook).Methods("POST")
}

func (s *Server) createBook(w http.ResponseWriter, r *http.Request) {
	var book application.CreateBookDTO
	switch r.Header.Get("Content-type") {
	case "application/json":
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			httpCode := errorStatusCode(application.ErrorCode(err))
			handleError(w, r, httpCode, application.ErrorMessage(err), err)
		}
	default:
		handleInvalidContentType(w, r)
	}
}
