package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"net/http"
)

type Server struct {
	Router *gin.Engine
	Mailer *sendgrid.Client
}

func NewServer(r *gin.Engine, m *sendgrid.Client) *Server {
	return &Server{
		Router: r,
		Mailer: m,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.Router.ServeHTTP(w, req)
}

func (s *Server) Run(addr string) error {
	log.Printf("Starting the server on: %v", addr)
	err := http.ListenAndServe(addr, s.routes())

	if err != nil {
		return err
	}

	return nil
}
