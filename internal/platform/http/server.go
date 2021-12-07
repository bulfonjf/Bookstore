package http

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
)

const ShutdownTimeout = 1 * time.Second

type Server struct {
	ln     net.Listener
	server *http.Server
	Router *mux.Router

	Addr   string
	Domain string

	RouteRegisters []RouteRegister
}

type RouteRegister interface {
	RegisterRoutes(router *mux.Router)
}

func NewServer() *Server {
	s := &Server{
		server: &http.Server{},
		Router: mux.NewRouter(),
	}

	return s
}

func (s *Server) Open() (err error) {
	s.Router.Use(reportPanic)

	s.Router.NotFoundHandler = http.HandlerFunc(HandleNotFound)

	router := s.Router.PathPrefix("/").Subrouter()
	for _, r := range s.RouteRegisters {
		r.RegisterRoutes(router)
	}

	s.server.Handler = router

	if s.Domain != "" {
		s.ln = autocert.NewListener(s.Domain)
	} else {
		if s.ln, err = net.Listen("tcp", s.Addr); err != nil {
			return err
		}
	}

	go s.server.Serve(s.ln)

	return nil
}

func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}

func reportPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func HandleInvalidContentType(w http.ResponseWriter, r *http.Request) {
	HandleErrorAsJson(w, r, http.StatusUnsupportedMediaType, "Content-type header must be application/json", nil)
}

func HandleNotAcceptable(w http.ResponseWriter, r *http.Request) {
	HandleErrorAsJson(w, r, http.StatusNotAcceptable, "The client doesn't accept application/json", nil)
}

func HandleBadScheme(w http.ResponseWriter, r *http.Request, err error) {
	HandleErrorAsJson(w, r, http.StatusBadRequest, "Invalid scheme", err)
}
