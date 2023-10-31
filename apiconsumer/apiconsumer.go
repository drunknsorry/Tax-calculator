package apiconsumer

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/drunknsorry/Tax-calculator/logger"
)

// Initiate logger
var logagr = logger.LoggerInit("log/api.log", "api: ")

// Setting client
var client http.Client

// Api url
var ApiUrl = "http://0.0.0.0:5000/tax-calculator/tax-year/"

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
	var response TaxBracketResults
	url := ApiUrl + year
	err := GetJson(url, &response)
	if err != nil {
		logagr.Printf("Error fetching results: %s: %v", url, err)
		return nil, err
	}
	return &response, nil

}

// Fetch json from specified API, check if any connection error is experienced and return decoded json
func GetJson(url string, resp interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		logagr.Printf("Error making http request: %s: %v", url, err)
		return err
	}

	// defer r.Body so it closes io.ReadCloser used internally by http.Response when GetJson exits regardless of result
	defer r.Body.Close()

	// Check for non 200 HTTP status codes
	if r.StatusCode == http.StatusNotFound {
		logagr.Printf("%s does not exist", url)
		return errors.New("HTTP 404 - Year data not found for your request, please try an accepted value")
	} else if r.StatusCode == http.StatusInternalServerError {
		logagr.Printf("%s experienced an internal server error", url)
		return errors.New("HTTP 500 - We've encountered an error, please try again later")
	} else if r.StatusCode != http.StatusOK {
		logagr.Printf("%s experienced an error", url)
		return errors.New("we've encountered an error, please try again later or contact us")
	}

	return json.NewDecoder(r.Body).Decode(resp)
}
