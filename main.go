package main

import (
	"fmt"
	"net/http"

	"github.com/drunknsorry/Tax-calculator/api"
	"github.com/drunknsorry/Tax-calculator/logger"
)

func main() {
	//Start a new mux instance
	mux := api.ServerStart()

	fmt.Println("Starting server on http://localhost:4000")
	logger.ServerLogger.Printf("Starting server on http://localhost:4000")

	serverError := http.ListenAndServe(":4000", mux)
	if serverError != nil {
		logger.ServerLogger.Printf("Server error: %v", serverError)
	}

}
