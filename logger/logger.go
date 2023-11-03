package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Add a logger
// var Logger *log.Logger

// Add an Api logger
var ApiLogger *log.Logger

// Add an Api Consumer logger
var ApiConsumerLogger *log.Logger

// Add a Server logger
var ServerLogger *log.Logger

func init() {

	currentTime := time.Now()
	fmt.Print(currentTime)

	dir, err := filepath.Abs("./log")
	if err != nil {
		log.Fatal("Error finding log folder or filepath")
	}

	// Check if directory exists, if not create it
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	// Hard coding each logger file, need to find a better implementation that's dynamic and scalable
	apiLogFile, apiErr := os.OpenFile(dir+"/api-"+currentTime.Format("2006-01-02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	apiConsumerLogFile, apiConsumerErr := os.OpenFile(dir+"/apiconsumer-"+currentTime.Format("2006-01-02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	serverLogFile, serverErr := os.OpenFile(dir+"/server-"+currentTime.Format("2006-01-02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if apiErr != nil || apiConsumerErr != nil || serverErr != nil {
		log.Fatal("Error opening log file:", apiErr, apiConsumerErr, serverErr)
	}

	ApiLogger = log.New(apiLogFile, "api: ", log.Ldate|log.Ltime|log.Lshortfile) // Using local day/time since it's easier rather than UTC
	ApiConsumerLogger = log.New(apiConsumerLogFile, "apiconsumer: ", log.Ldate|log.Ltime|log.Lshortfile)
	ServerLogger = log.New(serverLogFile, "server: ", log.Ldate|log.Ltime|log.Lshortfile)

}
