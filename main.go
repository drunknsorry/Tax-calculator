package main

import (
	"fmt"
	"net/http"

	"github.com/drunknsorry/Tax-calculator/api"
	"github.com/drunknsorry/Tax-calculator/logger"
)

// Initiate logger
var logagr = logger.LoggerInit("log/server.log", "server: ")

func main() {

	//Start a new mux instance
	mux := api.ServerStart()

	fmt.Println("Starting server on http://localhost:4000")
	logagr.Printf("Starting server on http://localhost:4000")

	serverError := http.ListenAndServe(":4000", mux)
	if serverError != nil {
		logagr.Printf("Server error: %v", serverError)
	}

}
