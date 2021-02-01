package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
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
	sendGridUser := os.Getenv("SEND_GRID_USER")
	sendGridPassword := os.Getenv("SEND_GRID_API_KEY")
	mailHost := "smtp.sendgrid.net"
	mailPort := 465

	gin.SetMode(gin.ReleaseMode)

	if port == "" {
		port = "5000"
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	r.Use(cors.Default())

	m := gomail.NewDialer(mailHost, mailPort, sendGridUser, sendGridPassword)

	server := app.NewServer(r, m)

	err := server.Run(":" + port)

	if err != nil {
		return err
	}

	return nil
}
