package app

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"log"
	"net/http"
)

type Server struct {
	Router *gin.Engine
	Mailer *gomail.Dialer
}

func NewServer(r *gin.Engine, m *gomail.Dialer) *Server {
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
