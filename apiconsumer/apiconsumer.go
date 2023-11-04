package apiconsumer

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/drunknsorry/Tax-calculator/logger"
	"github.com/drunknsorry/Tax-calculator/models"
)

// Setting client
var client http.Client

// Api url
var ApiUrl = "http://0.0.0.0:5000/tax-calculator/tax-year/"

// Fetch results or error from URL
func FetchResults(year string) (*models.TaxBracketResults, error) {
	var response models.TaxBracketResults
	url := ApiUrl + year
	err := GetJson(url, &response)
	if err != nil {
		logger.ApiConsumerLogger.Printf("Error fetching results: %s: %v", url, err)
		return nil, err
	}
	return &response, nil

}

// Fetch json from specified API, check if any connection error is experienced and return decoded json
func GetJson(url string, resp interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		logger.ApiConsumerLogger.Printf("Error making http request: %s: %v", url, err)
		return err
	}

	// defer r.Body so it closes io.ReadCloser used internally by http.Response when GetJson exits regardless of result
	defer r.Body.Close()

	// Check for non 200 HTTP status codes
	if r.StatusCode == http.StatusNotFound {
		logger.ApiConsumerLogger.Printf("%s does not exist", url)
		return errors.New("HTTP 404 - Year data not found for your request, please try an accepted value")
	} else if r.StatusCode == http.StatusInternalServerError {
		logger.ApiConsumerLogger.Printf("%s experienced an internal server error", url)
		return errors.New("HTTP 500 - We've encountered an error, please try again later")
	} else if r.StatusCode != http.StatusOK {
		logger.ApiConsumerLogger.Printf("%s experienced an error", url)
		return errors.New("we've encountered an error, please try again later or contact us")
	}

	return json.NewDecoder(r.Body).Decode(resp)
}
