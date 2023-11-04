package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/drunknsorry/Tax-calculator/models"
)

// Create test for Calculate Tax
func TestCalculateTax(t *testing.T) {
	// Assign mock data
	data := &models.TaxBracketResults{
		TaxBrackets: []models.Brackets{
			{Min: 0, Max: 50000, Rate: 0.10},
			{Min: 50000, Max: 75000, Rate: 0.05},
			{Min: 75000, Max: 100000, Rate: 0.20},
		},
	}

	year := 2019
	totalSalary := 100000

	tax, taxesPerBand, salaryPerBand, effectiveTaxRate := CalculateTaxes(data, float64(year), float64(totalSalary))

	// Test output of taxes per band
	expectedTaxesPerBand := []float64{5000, 1250, 5000}
	for i := 0; i < len(taxesPerBand); i++ {
		if taxesPerBand[i] != expectedTaxesPerBand[i] {
			t.Errorf("Expected Taxes Owed Per Band: %f, Got: %f", expectedTaxesPerBand[i], taxesPerBand[i])
		}
	}

	// Test output of Salary per band
	expectedSalaryPerBand := []float64{50000, 25000, 25000}
	for i := 0; i < len(salaryPerBand); i++ {
		if salaryPerBand[i] != expectedSalaryPerBand[i] {
			t.Errorf("Expected Salary Amount Per Tax Band: %f, Got: %f", expectedSalaryPerBand[i], salaryPerBand[i])
		}
	}

	// Test output of effective tax rate
	expectedEffectiveTaxRate := 0.11
	if expectedEffectiveTaxRate != effectiveTaxRate {
		t.Errorf("Expected Effective Tax Rate: %f, Got: %f", expectedEffectiveTaxRate, effectiveTaxRate)
	}

	// Test output of total tax
	expectedTax := float64(11250)
	if expectedTax != tax {
		t.Errorf("Expected Total Tax: %f, Got: %f", expectedTax, tax)
	}

	os.RemoveAll("log")
}

func TestHome(t *testing.T) {
	// Request home
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create mock response recorder
	rr := httptest.NewRecorder()

	// start a test server and set handler
	handler := http.HandlerFunc(home)
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code: %v, got %v", http.StatusOK, status)
	}

	// Clean up any log folders created
	os.RemoveAll("log")
}
