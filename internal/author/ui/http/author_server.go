package http

import (
	"bookstore/internal/apperror"
	"bookstore/internal/author/application"
	platform_http "bookstore/internal/platform/http"
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
)

type AuthorServer struct {
	AuthorService *application.AuthorService
}

func NewAuthorServer(authorService *application.AuthorService) *AuthorServer {
	return &AuthorServer{
		AuthorService: authorService,
	}
}

func (as *AuthorServer) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/authors", as.getAuthors).Methods("GET")
	router.HandleFunc("/authors/{authorID}", as.getAuthor).Methods("GET")
	router.HandleFunc("/authors", as.createAuthor).Methods("POST")
	router.HandleFunc("/authors", as.updateAuthor).Methods("PUT")
	router.HandleFunc("/authors/{authorID}", as.deleteAuthor).Methods("DELETE")
}

func (as *AuthorServer) getAuthor(w http.ResponseWriter, r *http.Request) {
	AuthorID := mux.Vars(r)["AuthorID"]
	switch r.Header.Get("Accept") {
	case "application/json":
		if AuthorID == "" {
			platform_http.HandleErrorAsJson(w, r, http.StatusBadRequest, "path param AuthorID is required", nil)

			return
		}

		AuthorDTO, err := as.AuthorService.GetAuthorByID(AuthorID)
		if err != nil {
			httpCode := platform_http.ErrorStatusCode(apperror.ErrorCode(err))
			platform_http.HandleErrorAsJson(w, r, httpCode, apperror.ErrorMessage(err), err)

			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(AuthorDTO)
	default:
		platform_http.HandleNotAcceptable(w, r)

		return
	}
}

func (as *AuthorServer) getAuthors(w http.ResponseWriter, r *http.Request) {
	AuthorID := mux.Vars(r)["AuthorID"]
	switch r.Header.Get("Accept") {
	case "application/json":
		if AuthorID == "" {
			platform_http.HandleErrorAsJson(w, r, http.StatusBadRequest, "path param AuthorID is required", nil)

			return
		}

		AuthorDTO, err := as.AuthorService.GetAuthorByID(AuthorID)
		if err != nil {
			httpCode := platform_http.ErrorStatusCode(apperror.ErrorCode(err))
			platform_http.HandleErrorAsJson(w, r, httpCode, apperror.ErrorMessage(err), err)

			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(AuthorDTO)
	default:
		platform_http.HandleNotAcceptable(w, r)

		return
	}
}

func (as *AuthorServer) createAuthor(w http.ResponseWriter, r *http.Request) {
	var createAuthorDTO application.CreateAuthorDTO
	switch r.Header.Get("Content-type") {
	case "application/json":
		if err := json.NewDecoder(r.Body).Decode(&createAuthorDTO); err != nil {
			platform_http.HandleBadScheme(w, r, err)

			return
		}
	default:
		platform_http.HandleInvalidContentType(w, r)

		return
	}

	switch r.Header.Get("Accept") {
	case "application/json":
		AuthorDTO, err := as.AuthorService.CreateAuthor(createAuthorDTO)
		if err != nil {
			httpCode := platform_http.ErrorStatusCode(apperror.ErrorCode(err))
			platform_http.HandleErrorAsJson(w, r, httpCode, apperror.ErrorMessage(err), err)

			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(AuthorDTO)
	default:
		platform_http.HandleNotAcceptable(w, r)

		return
	}
}

func (as *AuthorServer) updateAuthor(w http.ResponseWriter, r *http.Request) {
	var updateAuthorDTO application.UpdateAuthorDTO
	switch r.Header.Get("Content-type") {
	case "application/json":
		if err := json.NewDecoder(r.Body).Decode(&updateAuthorDTO); err != nil {
			platform_http.HandleBadScheme(w, r, err)

			return
		}
	default:
		platform_http.HandleInvalidContentType(w, r)

		return
	}

	switch r.Header.Get("Accept") {
	case "application/json":
		AuthorDTO, err := as.AuthorService.UpdateAuthor(updateAuthorDTO)
		if err != nil {
			httpCode := platform_http.ErrorStatusCode(apperror.ErrorCode(err))
			platform_http.HandleErrorAsJson(w, r, httpCode, apperror.ErrorMessage(err), err)

			return
		}

		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(AuthorDTO)

		return
	default:
		platform_http.HandleNotAcceptable(w, r)

		return
	}
}

func (as *AuthorServer) deleteAuthor(w http.ResponseWriter, r *http.Request) {
	AuthorID := mux.Vars(r)["AuthorID"]
	switch r.Header.Get("Accept") {
	case "application/json":
		if AuthorID == "" {
			platform_http.HandleErrorAsJson(w, r, http.StatusBadRequest, "path param AuthorID is required", nil)

			return
		}

		err := as.AuthorService.DeleteAuthor(AuthorID)
		if err != nil {
			httpCode := platform_http.ErrorStatusCode(apperror.ErrorCode(err))
			platform_http.HandleErrorAsJson(w, r, httpCode, apperror.ErrorMessage(err), err)

			return
		}

		w.WriteHeader(http.StatusNoContent)

		return
	default:
		platform_http.HandleNotAcceptable(w, r)

		return
	}
}
