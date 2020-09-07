package pkg

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

func (s *Server) TestMail() gin.HandlerFunc {
	return func(c *gin.Context) {
		mail := gomail.NewMessage()
		mail.SetAddressHeader("From", "hello@mbvistisen.dk", "Vis-ti-sen")
		mail.SetHeader("To", "vistisen@live.dk")
		mail.SetHeader("Subject", "Hej fra server")

		err := s.Mailer.DialAndSend(mail)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status":err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "great success"})
	}
}
