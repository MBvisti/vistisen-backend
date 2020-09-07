package app

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
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
	Name string `json:"name"`
	Mail string `json:"mail"`
	Subject string `json:"subject"`
	Msg string `json:"message"`
}

func (s *Server) TestMail() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var cM ContactMail

		err := c.ShouldBindJSON(cM)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": err})
			return
		}

		mail := gomail.NewMessage()
		mail.SetAddressHeader("From", cM.Mail, cM.Name)
		mail.SetHeader("To", "vistisen@live.dk")
		mail.SetHeader("Subject", cM.Subject)
		mail.SetBody("text/html", cM.Msg)

		err = s.Mailer.DialAndSend(mail)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
