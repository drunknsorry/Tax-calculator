package api

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/drunknsorry/Tax-calculator/apiconsumer"
)

// Add a logger
var logger *log.Logger

// Start the logger, create, open or append data to end of file, chmod for read write to owner and group
func loggerInit() {
	logFile, err := os.OpenFile("log/api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	logger = log.New(logFile, "api: ", log.Ldate|log.Ltime|log.Lshortfile) // Using local day/time since it's easier rather than UTC
}

// A function to instantiate the server
func ServerStart() *http.ServeMux {
	loggerInit()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/gettax", routeGetTax)
	return mux
}

// Define home route
func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome to the Tax Calculator API. Visit https://documenter.getpostman.com/view/30818865/2s9YXbARSe for Api documentation.")
}

// Define getTax route and it's logic
func routeGetTax(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	totalSalaryStr := r.URL.Query().Get("salary")

	// Check if year and salary values are empty, return error if they are
	if yearStr == "" {
		http.Error(w, "Year value not provided", http.StatusBadRequest)
		return
	}

	if totalSalaryStr == "" {
		http.Error(w, "Salary value not provided", http.StatusBadRequest)
		return
	}

	// Check if year or salary values are numbers that can be used to calculate taxes, return error if not
	totalSalary, errSalary := strconv.ParseFloat(totalSalaryStr, 64)

	if errSalary != nil {
		http.Error(w, "Invalid salary value", http.StatusBadRequest)
		return
	}

	year, errYear := strconv.ParseFloat(yearStr, 64)

	if errYear != nil {
		http.Error(w, "Invalid year value", http.StatusBadRequest)
		return
	}

	// Fetch data from api, if it fails return error
	data, err := apiconsumer.FetchResults(yearStr)
	if err != nil {
		logger.Printf("Failed fetching tax brackets: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Calculate taxes and assign to variables so a map can be created
	tax, taxesPerBand, salaryPerBand, effectiveTaxRate := CalculateTaxes(data, float64(year), float64(totalSalary))

	response := map[string]interface{}{
		"Salary":              totalSalary,
		"Total Taxes Owed":    tax,
		"Taxes Per Band":      taxesPerBand,
		"Salary Per Tax Band": salaryPerBand,
		"Effective Tax Rate":  effectiveTaxRate,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// Calculate taxes and return output data, round floats to 2 digits
func CalculateTaxes(data *apiconsumer.TaxBracketResults, year, totalSalary float64) (float64, []float64, []float64, float64) {
	tax := 0.0
	taxesPerBand := make([]float64, len(data.TaxBrackets))
	salaryPerBand := make([]float64, len(data.TaxBrackets))

	// Loop through slice of TaxBrackets
	for i, bracket := range data.TaxBrackets {
		if totalSalary <= bracket.Min { // Continue to next loop if salary is lower than min
			continue
		}
		if totalSalary >= bracket.Max && bracket.Max != 0.0 { // If salary is greater than max and max is not zero, do (max - min) * bracket tax rate
			salaryPerBand[i] = (bracket.Max - bracket.Min)
			taxesPerBand[i] = math.Round(((bracket.Max-bracket.Min)*bracket.Rate)*100) / 100
		} else {
			salaryPerBand[i] = (totalSalary - bracket.Min)
			taxesPerBand[i] = math.Round(((totalSalary-bracket.Min)*bracket.Rate)*100) / 100 // If salary is less than max, do (salary - min) * bracket tax rate
		}
		tax += taxesPerBand[i]
	}

	effectiveTaxRate := math.Round((tax/totalSalary)*100) / 100
	tax = math.Round(tax*100) / 100
	return tax, taxesPerBand, salaryPerBand, effectiveTaxRate
}