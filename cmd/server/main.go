package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

	if port == "" {
		port = "5000"
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	server := pkg.NewServer(r)

	err := server.Run(":" + port)

	if err != nil {
		return err
	}

	return nil
}
