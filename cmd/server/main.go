package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"os"
	"vistisen-backend/pkg"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	// setup dependencies
	port := os.Getenv("PORT")
	gmailAcc := os.Getenv("GMAIL_ACCOUNT")
	gmailPass := os.Getenv("GMAIL_PASSWORD")
	mailHost := os.Getenv("HOST")

	if port == "" {
		port = "5000"
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	m := gomail.NewDialer(mailHost, 465, gmailAcc, gmailPass)

	server := pkg.NewServer(r, m)

	err := server.Run(":" + port)

	if err != nil {
		return err
	}

	return nil
}
