package apiconsumer

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

// Add a logger
var logger *log.Logger

// Setting client
var client http.Client

// Api url
var ApiUrl = "http://0.0.0.0:5000/tax-calculator/tax-year/"

// Start the logger, create, open or append data to end of file, chmod for read write to owner and group
func loggerInit() {
	logFile, err := os.OpenFile("log/apiconsumer.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	logger = log.New(logFile, "apiconsumer: ", log.Ldate|log.Ltime|log.Lshortfile) // Using local time since it's easier rather than UTC
}

// Struct defining the data structure for the json payload
type TaxBracketResults struct {
	TaxBrackets []Brackets `json:"tax_brackets"`
}

// Struct defining the slice inside of TaxBrackets from TaxBracketResults
type Brackets struct {
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Rate float64 `json:"rate"`
}

// Fetch results or error from URL
func FetchResults(year string) (*TaxBracketResults, error) {
	loggerInit()
	var response TaxBracketResults
	url := ApiUrl + year
	err := GetJson(url, &response)
	if err != nil {
		logger.Printf("Error fetching results: %s: %v", url, err)
		return nil, err
	}
	return &response, nil

}

// Fetch json from specified API, check if any connection error is experienced and return decoded json
func GetJson(url string, resp interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		logger.Printf("Error making http request: %s: %v", url, err)
		return err
	}

	// defer r.Body so it closes io.ReadCloser used internally by http.Response when GetJson exits regardless of result
	defer r.Body.Close()

	// Check for non 200 HTTP status codes
	if r.StatusCode == http.StatusNotFound {
		logger.Printf("%s does not exist", url)
		return errors.New("HTTP 404 - Year data not found for your request, please try an accepted value")
	} else if r.StatusCode == http.StatusInternalServerError {
		logger.Printf("%s experienced an internal server error", url)
		return errors.New("HTTP 500 - We've encountered an error, please try again later")
	} else if r.StatusCode != http.StatusOK {
		logger.Printf("%s experienced an error", url)
		return errors.New("we've encountered an error, please try again later or contact us")
	}

	return json.NewDecoder(r.Body).Decode(resp)
}
