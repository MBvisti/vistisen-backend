package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	router *mux.Router
}

func NewServer(r *mux.Router) *Server {
	return &Server{
		router: r,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func (s *Server) Run(addr string) error {
	s.routes()

	log.Printf("Starting the server on: %v", addr)
	err := http.ListenAndServe(addr, s.router)

	if err != nil {
		return err
	}

	return nil
}

func (s *Server) ApiStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("vistisen backend api is good to go")
	}
}
