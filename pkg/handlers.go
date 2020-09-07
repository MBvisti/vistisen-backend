package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/http"
)

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

type ContactMail struct {
	Name    string `json:"name"`
	Mail    string `json:"mail"`
	Subject string `json:"subject"`
	Msg     string `json:"message"`
}

func (s *Server) Contact() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var contactMail ContactMail

		err := c.ShouldBindJSON(&contactMail)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}

		from := mail.NewEmail(contactMail.Name, contactMail.Mail)
		subject := contactMail.Subject
		to := mail.NewEmail("vis-ti-sen", "vistisen@live.dk")
		body := contactMail.Msg

		m := mail.NewSingleEmail(from, subject, to, body, "")
		response, err := s.Mailer.Send(m)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": response})
	}
}
