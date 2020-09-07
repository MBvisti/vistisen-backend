package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
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
	sendgridAPI := os.Getenv("SEND_GRID_API_KEY")
	gin.SetMode(gin.ReleaseMode)

	if port == "" {
		port = "5000"
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	r.Use(cors.Default())

	mailClient := sendgrid.NewSendClient(sendgridAPI)

	server := app.NewServer(r, mailClient)

	err := server.Run(":" + port)

	if err != nil {
		return err
	}

	return nil
}
