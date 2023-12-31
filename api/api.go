package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/drunknsorry/Tax-calculator/apiconsumer"
	"github.com/drunknsorry/Tax-calculator/logger"
	"github.com/drunknsorry/Tax-calculator/models"
	"github.com/drunknsorry/Tax-calculator/taxes"
)

// A function to instantiate the server
func ServerStart() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/gettax", routeGetTax)
	return mux
}

// Define home route
func home(w http.ResponseWriter, r *http.Request) {
	logger.ApiLogger.Printf("requested: %v", r.RequestURI)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome to the Tax Calculator API. Visit https://documenter.getpostman.com/view/30818865/2s9YXbARSe for Api documentation.")
}

// Define getTax route and it's logic
func routeGetTax(w http.ResponseWriter, r *http.Request) {
	logger.ApiLogger.Printf("requested: %v", r.RequestURI)

	// Reject all unsupported methods
	if r.Method != http.MethodGet {
		logger.ApiLogger.Printf("Unsupported Method: %v", r.Method)
		http.Error(w, "Unsupported Method", http.StatusBadRequest)
		return
	}

	yearStr := r.URL.Query().Get("year")
	totalSalaryStr := r.URL.Query().Get("salary")

	// Check if year and salary values are empty, return error if they are
	if yearStr == "" {
		logger.ApiLogger.Printf("Year value not provided %v", yearStr)
		http.Error(w, "Year value not provided", http.StatusBadRequest)
		return
	}

	if totalSalaryStr == "" {
		logger.ApiLogger.Printf("Salary value not provided %v", totalSalaryStr)
		http.Error(w, "Salary value not provided", http.StatusBadRequest)
		return
	}

	// Check if year or salary values are numbers that can be used to calculate taxes, return error if not
	totalSalary, errSalary := strconv.ParseFloat(totalSalaryStr, 64)

	if errSalary != nil {
		logger.ApiLogger.Printf("Invalid salary value: %v", totalSalaryStr)
		http.Error(w, "Invalid salary value", http.StatusBadRequest)
		return
	}

	year, errYear := strconv.ParseFloat(yearStr, 64)

	if errYear != nil || year < 2019 || year > 2022 {
		logger.ApiLogger.Printf("Invalid year value: %v", yearStr)
		http.Error(w, "Invalid year value", http.StatusBadRequest)
		return
	}

	// Fetch data from api, if it fails return error
	data, err := apiconsumer.FetchResults(yearStr)
	if err != nil {
		logger.ApiLogger.Printf("Failed fetching tax brackets: %v", err)
		http.Error(w, "Failed fetching tax brackets, please try again later", http.StatusInternalServerError)
		return
	}

	// Calculate taxes and assign to variables so a map can be created
	tax, taxesPerBand, salaryPerBand, effectiveTaxRate := taxes.CalculateTaxes(data, float64(year), float64(totalSalary))

	// Model response into a struct to pass on to json encoder
	response := models.TaxResponse{
		Salary:           totalSalary,
		TotalTaxesOwed:   tax,
		TaxesPerBand:     taxesPerBand,
		SalaryPerTaxBand: salaryPerBand,
		EffectiveTaxRate: effectiveTaxRate,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
