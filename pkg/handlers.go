package app

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"html/template"
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

		var cM ContactMail

		err := c.ShouldBindJSON(&cM)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "bad request from client", "res_code": 400})
			return
		}

		mailTemplate, err := template.ParseFiles("./contact_mail.html")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "server side error", "res_code": 500})
		}

		var t bytes.Buffer
		err = mailTemplate.Execute(&t, cM)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "server side error", "res_code": 500})
		}

		mail := gomail.NewMessage()
		mail.SetAddressHeader("From", "noreply@mbvistisen.dk", cM.Name+" wants to contact you")
		mail.SetHeader("To", "morten@mbvistisen.dk")
		mail.SetHeader("Subject", cM.Subject)
		mail.SetBody("text/html", t.String())

		err = s.Mailer.DialAndSend(mail)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "server side error", "res_code": 500})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "res_code": 200})
	}
}
