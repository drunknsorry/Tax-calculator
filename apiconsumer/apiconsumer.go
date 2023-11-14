package apiconsumer

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/drunknsorry/Tax-calculator/logger"
	"github.com/drunknsorry/Tax-calculator/models"
)

// Setting client
var client http.Client

// Api url
var ApiUrl = "http://0.0.0.0:5000/tax-calculator/tax-year/"

// Known data retrieval start and end years
var startYear = 2019
var endYear = 2022

// Create a map to "cache" responses
var ConsumerRequestCache = make(map[string]*models.TaxBracketResults)

// Fetch results or error from URL
func FetchResults(year string) (*models.TaxBracketResults, error) {

	// Check if values are prefetched and return for faster responses
	if ConsumerRequestCache[year] != nil {
		return ConsumerRequestCache[year], nil
	}

	var response models.TaxBracketResults
	url := ApiUrl + year
	err := GetJson(url, &response)
	if err != nil {
		logger.ApiConsumerLogger.Printf("Error fetching results: %s: %v", url, err)
		return nil, err
	}
	// Since the data was not found earlier, we are saving it into the cache
	ConsumerRequestCache[year] = &response
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

// Fetch results in the background, three attempts will be made.
func backgroundSync() {
	logger.ApiConsumerLogger.Printf("Starting background fetch of known year values")
	for i := startYear; i <= endYear; i++ {
		logger.ApiConsumerLogger.Printf("Fetching year: %v", i)
		for attempt := 1; attempt <= 3; attempt++ {
			yearResult, err := FetchResults(strconv.Itoa(i))
			if err != nil {
				logger.ApiConsumerLogger.Printf("Error making http request: %s, attempt: %v", err, attempt)
			} else {
				ConsumerRequestCache[strconv.Itoa(i)] = yearResult
				break
			}
		}
	}
}

func init() {
	// Run as a go routine so server starts without waiting on background sync
	go backgroundSync()

}
