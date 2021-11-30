package http

import (
	"bookstore/application"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) registerBookRoutes(r *mux.Router) {
	r.HandleFunc("/books", s.getBooks).Methods("GET")
	r.HandleFunc("/books", s.createBook).Methods("POST")
}

func (s *Server) getBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get("Accept") {
	case "application/json":
		booksDTO, err := s.BookService.GetBooks()
		if err != nil {
			httpCode := errorStatusCode(application.ErrorCode(err))
			handleError(w, r, httpCode, application.ErrorMessage(err), err)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(booksDTO)
	default:
		handleNotAcceptable(w, r)
	}
}

func (s *Server) createBook(w http.ResponseWriter, r *http.Request) {
	var createBookDTO application.CreateBookDTO
	switch r.Header.Get("Content-type") {
	case "application/json":
		if err := json.NewDecoder(r.Body).Decode(&createBookDTO); err != nil {
			handleBadScheme(w, r, err)
		}
	default:
		handleInvalidContentType(w, r)
	}

	switch r.Header.Get("Accept") {
	case "application/json":
		bookDTO, err := s.BookService.CreateBook(createBookDTO)
		if err != nil {
			httpCode := errorStatusCode(application.ErrorCode(err))
			handleError(w, r, httpCode, application.ErrorMessage(err), err)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(bookDTO)
	default:
		handleNotAcceptable(w, r)
	}
}
