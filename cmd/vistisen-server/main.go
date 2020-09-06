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

func run() error{
	// setup dependencies
	//buildEnv := os.Getenv("BUILD_ENV")
	//port := os.Getenv("port")
	r := mux.NewRouter()
	server := app.NewServer(r)

	err := server.Run(":8080")

	if err != nil {
		return err
	}

	return nil
}
