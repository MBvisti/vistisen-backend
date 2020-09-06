package pkg

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	Router *gin.Engine
}

func NewServer(r *gin.Engine) *Server {
	return &Server{
		Router: r,
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

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "vistisen backend api running smootly",
		}

		c.JSON(http.StatusOK, response)
	}
}
