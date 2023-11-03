package logger

import (
	"log"
	"os"
	"path/filepath"
)

// Add a logger
// var Logger *log.Logger

// Add an Api logger
var ApiLogger *log.Logger

// Add an Api Consumer logger
var ApiConsumerLogger *log.Logger

// Add a Server logger
var ServerLogger *log.Logger

// Start the logger, create, open or append data to end of file, chmod for read write to owner and group
// func LoggerInit(file string, logPrefix string) *log.Logger {
// 	logFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatal("Error opening log file:", err)
// 	}
// 	Logger = log.New(logFile, "api: ", log.Ldate|log.Ltime|log.Lshortfile) // Using local day/time since it's easier rather than UTC
// 	return Logger

// }

func init() {
	dir, err := filepath.Abs("../log")
	if err != nil {
		log.Fatal("Error finding log folder or filepath")
	}

	// Hard coding each logger file, need to find a better implementation that's dynamic and scalable
	apiLogFile, apiErr := os.OpenFile(dir+"/api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	apiConsumerLogFile, apiConsumerErr := os.OpenFile(dir+"/apiconsumer.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	serverLogFile, serverErr := os.OpenFile(dir+"/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if apiErr != nil || apiConsumerErr != nil || serverErr != nil {
		log.Fatal("Error opening log file:", apiErr, apiConsumerErr, serverErr)
	}

	ApiLogger = log.New(apiLogFile, "api: ", log.Ldate|log.Ltime|log.Lshortfile) // Using local day/time since it's easier rather than UTC
	ApiConsumerLogger = log.New(apiConsumerLogFile, "apiconsumer: ", log.Ldate|log.Ltime|log.Lshortfile)
	ServerLogger = log.New(serverLogFile, "server: ", log.Ldate|log.Ltime|log.Lshortfile)

}
