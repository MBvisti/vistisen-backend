package app

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
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
			log.Printf("this is the err: %v on line: %v", err, 42)
			c.JSON(http.StatusBadRequest, gin.H{"status": "bad request from client", "res_code": 400})
			return
		}

		mailTemplate, err := template.ParseFiles("./pkg/contact_mail.html")

		if err != nil {
			log.Printf("this is the err: %v on line: %v", err, 49)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "server side error", "res_code": 500})
		}

		var t bytes.Buffer
		err = mailTemplate.Execute(&t, cM)

		if err != nil {
			log.Printf("this is the err: %v on line: %v", err, 56)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "server side error", "res_code": 500})
		}

		mail := gomail.NewMessage()
		mail.SetAddressHeader("From", cM.Mail, cM.Name+" wants to contact you")
		mail.SetHeader("To", "heymbv@gmail.com")
		mail.SetHeader("Subject", cM.Subject)
		mail.SetBody("text/html", t.String())

		err = s.Mailer.DialAndSend(mail)

		if err != nil {
			log.Printf("this is the err: %v on line: %v", err, 67)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "server side error", "res_code": 500})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "res_code": 200})
	}
}
