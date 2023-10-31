package logger

import (
	"log"
	"os"
)

// Add a logger
var Logger *log.Logger

// Start the logger, create, open or append data to end of file, chmod for read write to owner and group
func LoggerInit(filePath string, logPrefix string) *log.Logger {
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	Logger = log.New(logFile, "api: ", log.Ldate|log.Ltime|log.Lshortfile) // Using local day/time since it's easier rather than UTC
	return Logger

}
