package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/drunknsorry/Tax-calculator/api"
)

// Add a logger
var logger *log.Logger

// Start the logger, create, open or append data to end of file, chmod for read write to owner and group
func loggerInit() {
	logFile, err := os.OpenFile("log/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	logger = log.New(logFile, "server: ", log.Ldate|log.Ltime|log.Lshortfile) // Using local day/time since it's easier rather than UTC
}

func main() {

	//Start a new mux instance
	mux := api.ServerStart()

	loggerInit()

	fmt.Println("Starting server on http://localhost:4000")
	logger.Printf("Starting server on http://localhost:4000")

	serverError := http.ListenAndServe(":4000", mux)
	if serverError != nil {
		logger.Printf("Server error: %v", serverError)
	}

}
