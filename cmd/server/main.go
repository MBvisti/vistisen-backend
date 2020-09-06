package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"os"
	app "vistisen-backend/pkg/server"
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
	}
	r := mux.NewRouter()
	server := app.NewServer(r)

	err := server.Run(":" + port)

	if err != nil {
		return err
	}

	return nil
}
